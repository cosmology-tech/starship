package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
func AddOrUpdateChartRepo() error {
	// Get the repo entry
	repoEntry := NewRepoEntry()

	repoFile, lock, err := CreateOrGetRepoFile(settings.RepositoryConfig)
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Unlock()
	}

	// if no repo entry exists, add it
	var chartRepo *repo.ChartRepository
	if !repoFile.Has(repoEntry.Name) {
		chartRepo, err = repo.NewChartRepository(repoEntry, getter.All(settings))
		if err != nil {
			return err
		}
	}



	return nil
}

func CreateOrGetRepoFile(repoFile string) (*repo.File, *flock.Flock, error) {
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

// fetchChart fetches a helm chart from a repo for a version
func fetchChart(chart, version string) (*repo.ChartRepository, error) {
	repoFile := settings.RepositoryConfig

	//Ensure the file directory exists as it is required for file locking
	err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		zap.L().Fatal("unable to create repo file", zap.Error(err))
	}

	// Acquire a file lock for process synchronization
	fileLock := flock.New(strings.Replace(repoFile, filepath.Ext(repoFile), ".lock", 1))
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
		defer fileLock.Unlock()
	}
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(repoFile)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	var f repo.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return nil, err
	}

	if f.Has(chart) {
		fmt.Printf("repository name (%s) already exists\n", chart)
		return nil, nil
	}

	entry := &repo.Entry{
		Name: chart,
		URL:  defaultHelmRepoURL,
	}

	r, err := repo.NewChartRepository(entry, getter.All(settings))
	if err != nil {
		return nil, err
	}

	return r, nil
}

func updateChart(chart string) error {
	repoFile := settings.RepositoryConfig

	f, err := repo.LoadFile(repoFile)
	if os.IsNotExist(errors.Cause(err)) || len(f.Repositories) == 0 {
		return errors.New("no repositories found. You must add one before updating"))
	}

	var repoChart *repo.ChartRepository

	for _, cfg := range f.Repositories {
		if cfg.Name == chart {
			r, err := repo.NewChartRepository(cfg, getter.All(settings))
			if err != nil {
				return err
			}
			repoChart = r
			break
		}
	}

	if repoChart == nil {
		// todo: fetchChart here

	}

	zap.L().Debug("hang tight while we grab the latest from your chart repositories")
	if _, err := repoChart.DownloadIndexFile(); err != nil {
		return fmt.Errorf("unable to get an update from the %q chart repository (%s), err :%s", repoChart.Config.Name, repoChart.Config.URL, err)
	}

	zap.L().Info(fmt.Sprintf("successfully got an update from the %q chart repository", repoChart.Config.Name))

	return nil
}
