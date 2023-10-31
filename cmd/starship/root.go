package main

import (
	"fmt"
	"github.com/cosmology-tech/starship/pkg/defaults"
	"github.com/cosmology-tech/starship/pkg/loader/starship"
	"github.com/cosmology-tech/starship/pkg/transformer/kubernetes"
	"github.com/cosmology-tech/starship/pkg/types"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func newStartCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:      "start",
		Usage:     "start starship resources from a config file",
		UsageText: "start [path to config-file] [options]",
		Flags:     GetCommandLineOptions(),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}
			// set configfile to the Config struct from args if not set
			if config.ConfigFile == "" {
				if c.NArg() > 0 {
					config.ConfigFile = c.Args().Get(0)
				} else {
					return cli.Exit("config file need to be specified", 1)
				}
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
	return &cli.Command{
		Name:  "list",
		Usage: "list starship charts deployed",
		Flags: GetCommandLineOptions("verbose"),
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
		Flags: GetCommandLineOptions("name", "version", "verbose"),
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

func newGenerateCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:      "generate",
		Usage:     "generate will generate yaml files from given starship config file",
		UsageText: "generate [path to config-file] [options]",
		Flags:     GetCommandLineOptions("config", "verbose"),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}
			// set configfile to the Config struct from args if not set
			if config.ConfigFile == "" {
				if c.NArg() > 0 {
					config.ConfigFile = c.Args().Get(0)
				} else {
					return cli.Exit("config file need to be specified", 1)
				}
			}
			log, err := NewLogger(config)
			if err != nil {
				return cli.Exit(err, 1)
			}

			// default config
			defaultConfig := defaults.DefaultConfig()

			// convert config file to nodeConfig from loader
			s := starship.NewStarship(log)
			nodeConfigs, err := s.LoadFile([]string{config.ConfigFile}, defaultConfig)
			if err != nil {
				return cli.Exit(fmt.Sprintf("unable to load config to nodeconfigs, err: %s", err), 1)
			}
			log.Debug("node config", zap.Any("node-configs", nodeConfigs))

			// convert nodeconfig to k8s objects
			k := kubernetes.NewKubernetes(log)
			objects, err := k.Transform(nodeConfigs, types.ConvertOptions{})
			if err != nil {
				return cli.Exit(fmt.Sprintf("error transforming nodeconfig to k8s objects, err: %s", err), 1)
			}
			log.Debug("kubernetes runtime objects", zap.Any("k8s", objects))

			err = objects.WriteToFile("build/generator/")
			if err != nil {
				return cli.Exit(err, 1)
			}

			return nil
		},
	}
}

func newConnectCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:      "connect",
		Usage:     "connect will perform port-forward based on the configfile to localhost ports",
		UsageText: "connect [path to config-file] [options]",
		Flags:     GetCommandLineOptions("config", "namespace", "verbose"),
		Action: func(c *cli.Context) error {
			if err := ParseCLIOptions(c, config); err != nil {
				return cli.Exit(err, 1)
			}
			// set configfile to the Config struct from args if not set
			if config.ConfigFile == "" {
				if c.NArg() > 0 {
					config.ConfigFile = c.Args().Get(0)
				} else {
					return cli.Exit("config file need to be specified", 1)
				}
			}

			client, err := NewClient(config)
			if err != nil {
				return cli.Exit(err, 1)
			}
			defer client.logger.Sync()

			// check kubectl is installed
			err = client.CheckKubectl()
			if err != nil {
				client.logger.Error(err.Error())
				return cli.Exit(err, 1)
			}

			err = client.RunPortForward(c.Context)
			if err != nil {
				client.logger.Error(err.Error())
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
		newConnectCommand(config),
		newGenerateCommand(config),
	}

	return app
}
