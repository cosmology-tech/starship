package main

import (
	"os"

	"helm.sh/helm/v3/pkg/cli"
)

var settings *cli.EnvSettings

const (
	// defaultRepositoryURL is the default URL of the Helm Hub repository.
	defaultHelmRepoURL = "https://cosmology-tech.github.io/starship/"
)

func main() {
	// Initialize settings variable
	settings = cli.New()
	settings.KubeConfig = os.Getenv("KUBECONFIG")

	chartVersion := "0.1.23"
	noWait := false
	cofigFile := "tests/configs/one-chain.yaml"

	// Fetch helm chart
	err := AddOrUpdateChartRepo(chartVersion)
	if err != nil {
		panic(err)
	}

	// Install helm chart, and wait for it to be ready
	err = InstallChart(chartVersion, cofigFile, !noWait)
	if err != nil {
		panic(err)
	}
}
