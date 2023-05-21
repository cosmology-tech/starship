package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

/*
Charts functions are supposed to
* If not present, fetch and upgrade the chart
* If installed, upgrade the chart
* Make sure the specified version is installed, if not upgrade the chart

Inputs:
* version, default will be set to latest
*/

func NewRepoEntry() *repo.Entry {
	return &repo.Entry{
		Name: "starship",
		URL:  defaultHelmRepoURL,
	}
}

// AddOrUpdateChartRepo adds or updates a chart repo in place
// for the given version of the chart
func AddOrUpdateChartRepo(version string) error {
	// Get the repo entry
	repoEntry := NewRepoEntry()

	repoFile, lock, err := createOrGetRepoFile(settings.RepositoryConfig)
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Unlock()
	}

	chartRepo, err := repo.NewChartRepository(repoEntry, getter.All(settings))
	if err != nil {
		return err
	}

	chartRepo.CachePath = settings.RepositoryCache

	idx, err := chartRepo.DownloadIndexFile()
	if err != nil {
		return err
	}

	// Update the repo file with the new entry
	repoFile.Update(repoEntry)

	// Read the index file for the repository to get chart information and return chart URL
	repoIndex, err := repo.LoadIndexFile(idx)
	if err != nil {
		return err
	}

	// check if version is available for the chart
	_, err = repoIndex.Get(repoEntry.Name, version)
	if err != nil {
		return fmt.Errorf("chart version is invalid: %s", err)
	}

	return nil
}

func createOrGetRepoFile(repoFile string) (*repo.File, *flock.Flock, error) {
	// Check if the repo file exists, if not create it
	if _, err := os.Stat(repoFile); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
		if err != nil {
			return nil, nil, err
		}
	}

	// Acquire a file lock for process synchronization
	fileLock := flock.New(strings.Replace(repoFile, filepath.Ext(repoFile), ".lock", 1))
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err != nil {
		return nil, nil, err
	}
	if !locked {
		return nil, nil, fmt.Errorf("unable to lock %s", repoFile)
	}

	b, err := os.ReadFile(repoFile)
	if err != nil && !os.IsNotExist(err) {
		return nil, nil, err
	}

	var f *repo.File
	if err := yaml.Unmarshal(b, f); err != nil {
		return nil, nil, err
	}

	return f, fileLock, nil
}
