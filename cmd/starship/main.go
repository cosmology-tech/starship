package main

import (
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

	// Fetch helm chart
	err := AddOrUpdateChartRepo("0.1.23")
	if err != nil {
		panic(err)
	}
}
