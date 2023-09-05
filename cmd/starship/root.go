package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func newStartCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start starship resources from a config file",
		Flags: GetCommandLineOptions(),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}

			client, err := NewClient(config)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer client.logger.Sync()

			if err := client.AddOrUpdateChartRepo(); err != nil {
				return cli.Exit(err, 1)
			}

			err = client.InstallChart()
			if err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}
}

func newListCommand(config *Config) *cli.Command {
	fmt.Printf("here....")
	return &cli.Command{
		Name:  "list",
		Usage: "list starship charts deployed",
		Flags: GetCommandLineOptions(),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}

			client, err := NewClient(config)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer client.logger.Sync()

			res, err := client.ListCharts()
			if err != nil {
				return cli.Exit(err, 1)
			}
			for _, r := range res {
				fmt.Printf("name: %s,\tstatus: %s,\tnamespace: %s\n", r.Name, r.Info.Status, r.Namespace)
			}

			return nil
		},
	}
}

func newStopCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop running starship resources",
		Flags: GetCommandLineOptions(),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}

			client, err := NewClient(config)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer client.logger.Sync()

			err = client.DeleteChart(config.Name)
			if err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}
}

func NewApp() *cli.App {
	config := NewDefaultConfig()
	app := cli.NewApp()
	app.Name = Prog
	app.Usage = Description
	app.Version = Version
	app.UsageText = "starship [options]"

	app.Action = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []*cli.Command{
		newStartCommand(config),
		newListCommand(config),
		newStopCommand(config),
	}

	return app
}
