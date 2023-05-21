package main

import (
	"os/exec"
)

/*
Charts functions are supposed to
* If not present, fetch and upgrade the chart
* If installed, upgrade the chart
* Make sure the specified version is installed, if not upgrade the chart
Inputs:
* version, default will be set to latest
*/

// AddOrUpdateChartRepo adds or updates a chart repo in place
// for the given version of the chart
func (c *Client) AddOrUpdateChartRepo() error {
	err := exec.Command("helm", "repo", "add", "starship", defaultHelmRepoURL).Run()
	if err != nil {
		return err
	}

	return nil
}

// InstallChart installs a chart
func (c *Client) InstallChart() error {
	return nil
}
