package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
	"sigs.k8s.io/yaml"
)

/*
Charts functions are supposed to
* If not present, fetch and upgrade the chart
* If installed, upgrade the chart
* Make sure the specified version is installed, if not upgrade the chart

Inputs:
* version, default will be set to latest
*/

func (c *Client) CreateRepoEntry() *repo.Entry {
	return &repo.Entry{
		Name: c.config.HelmChartName,
		URL:  c.config.HelmRepoURL,
	}
}

// AddOrUpdateChartRepo adds or updates a chart repo in place
// for the given version of the chart
func (c *Client) AddOrUpdateChartRepo() error {
	// Get the repo entry
	repoEntry := c.CreateRepoEntry()

	repoFile, lock, err := c.createOrGetRepoFile()
	if err != nil {
		return err
	}
	if lock != nil {
		defer lock.Unlock()
	}

	chartRepo, err := repo.NewChartRepository(repoEntry, getter.All(c.settings))
	if err != nil {
		return err
	}

	chartRepo.CachePath = c.settings.RepositoryCache

	idx, err := chartRepo.DownloadIndexFile()
	if err != nil {
		return err
	}

	// Update the repo file with the new entry
	//repoFile.Update(repoEntry)
	c.logger.Debug("repofile", zap.Any("repoFile", repoFile))
	if !repoFile.Has(repoEntry.Name) {
		c.logger.Info("repo file does not have entry", zap.String("name", repoEntry.Name))
	}

	// Read the index file for the repository to get chart information and return chart URL
	repoIndex, err := repo.LoadIndexFile(idx)
	if err != nil {
		return err
	}

	// check if version is available for the chart
	_, err = repoIndex.Get(repoEntry.Name, c.config.Version)
	if err != nil {
		return fmt.Errorf("chart version is invalid: %s", err)
	}

	return nil
}

func (c *Client) createOrGetRepoFile() (*repo.File, *flock.Flock, error) {
	repoFile := c.settings.RepositoryConfig

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

	var f repo.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return nil, nil, err
	}

	return &f, fileLock, nil
}

func (c *Client) createInstallClient(config *action.Configuration) *action.Install {
	client := action.NewInstall(config)
	client.Namespace = c.config.Namespace
	client.CreateNamespace = true
	client.ReleaseName = c.config.Name
	client.Version = c.config.Version
	client.Wait = c.config.Wait
	client.Timeout = 300 * time.Second
	client.Atomic = c.config.Atomic
	client.Force = false

	return client
}

// InstallChart installs a chart
func (c *Client) InstallChart() error {
	actionConfig := new(action.Configuration)
	err := actionConfig.Init(c.settings.RESTClientGetter(), c.settings.Namespace(), "configmap", func(format string, v ...interface{}) {
		c.logger.Debug(fmt.Sprintf(format, v...))
	})
	if err != nil {
		return err
	}

	// if chart exists, provide warning to user
	r, err := c.GetChart(c.config.Name)
	if err != nil {
		if err.Error() != "not found" {
			return nil
		}
	}
	if r != nil {
		c.logger.Error("chart already exists, please stop the current chart with `stop` commad before installing", zap.String("name", r.Name))
		return fmt.Errorf("chart already exists: name: %s", r.Name)
	}

	c.logger.Info("installing chart, hang tight...")
	client := c.createInstallClient(actionConfig)

	cp, err := client.ChartPathOptions.LocateChart(
		fmt.Sprintf("%s/%s", c.config.HelmRepoName, c.config.HelmChartName),
		c.settings,
	)
	if err != nil {
		return err
	}

	// Get all the values from the config file
	p := getter.All(c.settings)
	valueOpts := &values.Options{
		ValueFiles: []string{c.config.ConfigFile},
	}
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return err
	}

	chartReq, err := loader.Load(cp)
	if err != nil {
		return err
	}

	r, err = client.Run(chartReq, vals)
	if err != nil {
		return err
	}

	if r.Info.Status == release.StatusDeployed {
		c.logger.Info("successfully deployed resources", zap.String("release", r.Name), zap.Any("info", r.Info))
	}

	return nil
}

func (c *Client) createListClient(config *action.Configuration) *action.List {
	client := action.NewList(config)
	client.ByDate = true

	return client
}

func (c *Client) ListCharts() ([]*release.Release, error) {
	actionConfig := new(action.Configuration)
	err := actionConfig.Init(c.settings.RESTClientGetter(), "", "configmap", func(format string, v ...interface{}) {
		c.logger.Debug(fmt.Sprintf(format, v...))
	})
	if err != nil {
		return nil, err
	}

	client := c.createListClient(actionConfig)

	return client.Run()
}

// GetChart gets the chart by name
func (c *Client) GetChart(name string) (*release.Release, error) {
	actionConfig := new(action.Configuration)
	err := actionConfig.Init(c.settings.RESTClientGetter(), "", "configmap", func(format string, v ...interface{}) {
		c.logger.Debug(fmt.Sprintf(format, v...))
	})
	if err != nil {
		return nil, err
	}

	client := c.createListClient(actionConfig)

	client.SetStateMask()

	releases, err := client.Run()
	if err != nil {
		return nil, err
	}

	var r *release.Release
	for _, r = range releases {
		if r.Name == name {
			return r, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (c *Client) DeleteChart(name string) error {
	actionConfig := new(action.Configuration)
	err := actionConfig.Init(c.settings.RESTClientGetter(), c.settings.Namespace(), "configmap", func(format string, v ...interface{}) {
		c.logger.Debug(fmt.Sprintf(format, v...))
	})
	if err != nil {
		return err
	}

	client := action.NewUninstall(actionConfig)
	_, err = client.Run(name)
	if err != nil {
		return err
	}

	return nil
}
