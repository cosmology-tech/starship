package main

import (
	"fmt"
	"os/exec"

	"go.uber.org/zap"
)

const (
	defaultBinaryPath = "~/.local/bin"
)

// VerifyOrInstallDeps verifies all dependencies are installed
// or installs the dependencies if they are not installed.
// This will install following dependencies
// * helm
// * kubectl
// * kind
func (c *Client) VerifyOrInstallDeps() error {
	for _, bin := range []string{"helm", "kubectl", "kind"} {
		if c.verfiyBinary(bin) {
			c.logger.Info("Binary already installed", zap.String("binary", bin))
			continue
		}

		err := c.installBinary(bin)
		if err != nil {
			return err
		}

		c.logger.Info("Binary installed", zap.String("binary", bin))
	}

	return nil
}

func (c *Client) verfiyBinary(bin string) bool {
	// Check if helm is installed
	_, err := exec.LookPath(fmt.Sprintf("%s/%s", defaultBinaryPath, bin))
	if err == nil {
		return true
	}
	return false
}

func (c *Client) installBinary(bin string) error {
	c.logger.Info("Installing binary", zap.String("binary", bin), zap.String("path", defaultBinaryPath))
	c.logger.Info("Note: we will require sudo access to install binaries")

	var err error
	switch bin {
	case "helm":
		err = exec.Command("sudo", "curl", "-fsSL", "https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3", "|", "bash").Run()
	case "kubectl":
		version, verr := exec.Command("sudo", "curl", "-Ls", "https://dl.k8s.io/release/stable.txt").Output()
		if verr != nil {
			return verr
		}
		err = exec.Command("sudo", "curl", "-Lks",
			fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/amd64/kubectl", string(version)),
			">", defaultBinaryPath, "&&", "chmod", "+x", fmt.Sprintf("%s/kubectl", defaultBinaryPath)).Run()
	case "kind":
		err = exec.Command("sudo", "curl", "-Lo", fmt.Sprintf("%s/kind", defaultBinaryPath), "https://kind.sigs.k8s.io/dl/v0.11.1/kind-linux-amd64").Run()
	}
	if err != nil {
		return err
	}

	if !c.verfiyBinary(bin) {
		c.logger.Error("Binary not installed, something went wrong", zap.String("binary", bin))
		return fmt.Errorf("binary %s not installed", bin)
	}

	return nil
}
