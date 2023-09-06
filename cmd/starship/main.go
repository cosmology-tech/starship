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
	log.Debug(
		"starting Starship",
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
	settings.SetNamespace(config.Namespace)
	settings.Debug = config.Verbose
	client.settings = settings

	return client, nil
}

func main() {
	app := NewApp()
	_ = app.Run(os.Args)
}
