package main

import (
	"os"

	"go.uber.org/zap"
	"helm.sh/helm/v3/pkg/cli"
)

type Client struct {
	config *Config
	logger *zap.Logger

	settings *cli.EnvSettings
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

	// Set settings
	settings := cli.New()
	settings.KubeConfig = os.Getenv("KUBECONFIG")
	settings.SetNamespace("aws-starship")
	settings.Debug = config.Verbose
	client.settings = settings

	return client, nil
}

func main() {
	config := NewDefaultConfig()
	client, err := NewClient(config)

	// Fetch helm chart
	err = client.AddOrUpdateChartRepo()
	if err != nil {
		panic(err)
	}

	// Install helm chart, and wait for it to be ready
	err = client.InstallChart()
	if err != nil {
		panic(err)
	}
}
