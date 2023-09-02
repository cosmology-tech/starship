package main

import (
	"fmt"
	"reflect"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewDefaultConfig() *Config {
	return &Config{
		Namespace:     "default",
		HelmRepoName:  "starship",
		HelmChartName: "devnet",
		Verbose:       true,
		Wait:          true,
		Version:       "0.1.45",
		HelmRepoURL:   "https://cosmology-tech.github.io/starship/",
		ConfigFile:    "config.yaml",
	}
}

type Config struct {
	Namespace     string `name:"namespace" json:"namespace" usage:"kubernetes namespace for deployment, default: default"`
	ConfigFile    string `name:"config-file" json:"config_file" usage:"path to the config file"`
	Version       string `name:"version" json:"version" usage:"version of the helm chart"`
	HelmRepoURL   string `name:"helm-repo-url" json:"helm_repo_url" usage:"helm repo url"`
	HelmRepoName  string `name:"helm-repo-name" json:"helm_repo_name" usage:"helm repo name"`
	HelmChartName string `name:"helm-chart-name" json:"helm_chart_name" usage:"helm chart name"`
	Wait          bool   `name:"wait" json:"wait" usage:"wait for the helm chart to be ready"`
	Verbose       bool   `name:"verbose" json:"verbose" usage:"switch on debug / verbose logging"`
	OnlyFatalLog  bool   `name:"only-fatal-log" json:"only_fatal_log" usage:"used while running test"`
}

func GetCommandLineOptions() []cli.Flag {
	defaults := NewDefaultConfig()
	var flags []cli.Flag
	count := reflect.TypeOf(Config{}).NumField()
	for i := 0; i < count; i++ {
		field := reflect.TypeOf(Config{}).Field(i)
		usage, found := field.Tag.Lookup("usage")
		if !found {
			continue
		}
		envName := field.Tag.Get("env")
		if envName != "" {
			envName = envPrefix + envName
		}
		optName := field.Tag.Get("name")

		switch t := field.Type; t.Kind() {
		case reflect.Bool:
			dv := reflect.ValueOf(defaults).Elem().FieldByName(field.Name).Bool()
			msg := fmt.Sprintf("%s (default: %t)", usage, dv)
			flags = append(flags, &cli.BoolFlag{
				Name:    optName,
				Usage:   msg,
				EnvVars: []string{envName},
			})
		case reflect.String:
			defaultValue := reflect.ValueOf(defaults).Elem().FieldByName(field.Name).String()
			flags = append(flags, &cli.StringFlag{
				Name:    optName,
				Usage:   usage,
				EnvVars: []string{envName},
				Value:   defaultValue,
			})
		}
	}

	return flags
}

func ParseCLIOptions(cx *cli.Context, config *Config) (err error) {
	// iterate the Config and grab command line options via reflection
	count := reflect.TypeOf(config).Elem().NumField()
	for i := 0; i < count; i++ {
		field := reflect.TypeOf(config).Elem().Field(i)
		name := field.Tag.Get("name")

		if cx.IsSet(name) {
			switch field.Type.Kind() {
			case reflect.Bool:
				reflect.ValueOf(config).Elem().FieldByName(field.Name).SetBool(cx.Bool(name))
			case reflect.String:
				reflect.ValueOf(config).Elem().FieldByName(field.Name).SetString(cx.String(name))
			}
		}
	}
	return nil
}

func NewLogger(config *Config) (*zap.Logger, error) {
	c := zap.NewProductionConfig()
	c.DisableCaller = true
	c.Encoding = "console"
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	c.EncoderConfig.CallerKey = zapcore.OmitKey
	c.EncoderConfig.TimeKey = zapcore.OmitKey
	c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	if config.Verbose {
		c.DisableCaller = false
		c.Development = true
		c.DisableStacktrace = true // Disable stack trace for development
		c.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	if config.OnlyFatalLog {
		c.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	}

	log, err := c.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(log) // Set zap global logger
	return log, err
}
