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
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
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

func NewRepoEntry() *repo.Entry {
	return &repo.Entry{
		Name: "devnet",
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
	//repoFile.Update(repoEntry)
	if !repoFile.Has(repoEntry.Name) {
		zap.L().Info("repo file does not have entry", zap.String("name", repoEntry.Name))
	}

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

	fmt.Printf("repofile name: %s", repoFile)

	var f repo.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return nil, nil, err
	}

	return &f, fileLock, nil
}

func getKubeClient() (*rest.Config, error) {
	config := genericclioptions.NewConfigFlags(true)

	kubeConfig := os.Getenv("KUBECONFIG")
	config.KubeConfig = &kubeConfig

	restConfig, err := config.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	return restConfig, nil
}

// InstallChart installs a chart
func InstallChart(version, cofigFile string, wait bool) error {
	//restConfig, err := getKubeClient()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf(restConfig.String())
	//settings.KubeConfig = os.Getenv("KUBECONFIG")

	// create kube client and make sure it is reachable
	//kubeClient := kube.New(settings.RESTClientGetter())
	//err := kubeClient.IsReachable()
	//if err != nil {
	//	return err
	//}
	//kubeClient.Namespace = ""

	actionConfig := new(action.Configuration)
	err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), "memory", func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	})
	if err != nil {
		return err
	}

	client := action.NewInstall(actionConfig)
	client.Namespace = "aws-starship"
	client.ReleaseName = "devnet"
	client.Version = version
	client.Wait = wait

	cp, err := client.ChartPathOptions.LocateChart("starship/devnet", settings)
	if err != nil {
		return err
	}

	// Get all the values from the config file
	p := getter.All(settings)
	valueOpts := &values.Options{
		ValueFiles: []string{cofigFile},
	}
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return err
	}

	chartReq, err := loader.Load(cp)
	if err != nil {
		return err
	}

	_, err = client.Run(chartReq, vals)
	if err != nil {
		return err
	}

	return nil
}
