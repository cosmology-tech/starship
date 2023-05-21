package main

import (
	"go.uber.org/zap"
)

type Client struct {
	config *Config
	logger *zap.Logger
}

func NewClient(config *Config) (*Client, error) {
	log, err := NewLogger(config)
	if err != nil {
		return nil, err
	}
	log.Info(
		"Starting the service",
		zap.String("prog", Prog),
		zap.String("version", Version),
		zap.Any("config", config),
	)

	client := &Client{
		config: config,
		logger: log,
	}

	return client, nil
}

func main() {
	config := NewDefaultConfig()
	client, err := NewClient(config)

	err = client.VerifyOrInstallDeps()
	if err != nil {
		return
	}

	chartVersion := "0.1.23"
	noWait := false
	cofigFile := "tests/configs/one-chain.yaml"

	// Fetch helm chart
	err = client.AddOrUpdateChartRepo(chartVersion)
	if err != nil {
		panic(err)
	}

	// Install helm chart, and wait for it to be ready
	err = InstallChart(chartVersion, cofigFile, !noWait)
	if err != nil {
		panic(err)
	}
}
