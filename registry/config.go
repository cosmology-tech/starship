package main

import (
	"fmt"
	"reflect"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func NewDefaultConfig() *Config {
	return &Config{
		Host:            "0.0.0.0",
		HTTPPort:        "8080",
		GRPCPort:        "9090",
		ChainRegistry:   "chains/",
		ChainClientHome: "~/.lens",
	}
}

type Config struct {
	// Host is the interface to bind the HTTP service on
	Host string `name:"host" json:"host" env:"HOST" usage:"Host address to listen on"`
	// HTTPPort is the port for the http server
	HTTPPort string `name:"http-port" json:"http_port" env:"HTTP_PORT" usage:"Port for http server"`
	// GRPCPort is the port for the grpc server
	GRPCPort string `name:"grpc-port" json:"grpc_port" env:"GRPC_PORT" usage:"Port for gRPC server"`
	// ChainRegistry is full path to the directory containing various chain registry information
	ChainRegistry string `name:"chain-registry" json:"chain_registry" env:"CHAIN_REGISTRY" usage:"Path of chain registry files"`
	// ChainClientHome is the path the home directory for lens client
	ChainClientHome string `name:"chain-client-home" json:"chain_client_home" env:"CHAIN_CLIENT_HOME" usage:"Path to the home of lens client directory"`
	// ChainClientIDs is a comma seperated list of chain ids for various chains
	ChainClientIDs string `name:"chain-client-ids" json:"chain_client_ids" env:"CHAIN_CLIENT_IDS" usage:"Comma seperated list of chain ids for various chains"`
	// ChainClientRPCs is a comma seperated list of chain rpc address for various chains, used to create connections
	// Note: ChainClientRPCs is different from ChainAPIRPCs, as ChainClientRPCs is used for internal routing
	// whereas ChainAPIRPCs is used for /chains endpoint result
	ChainClientRPCs string `name:"chain-client-rpcs" json:"chain_client_rpcs" env:"CHAIN_CLIENT_RPCS" usage:"Comma seperated list of chain rpc address for various chains"`
	// ChainAPIRPCs is a comma seperated list of chain rpc address for various chains, used at output of chain.apis.rpc
	ChainAPIRPCs string `name:"chain-api-rpcs" json:"chain_api_rpcs" env:"CHAIN_API_RPCS" usage:"Comma seperated list of chain rpc address, used at output of chain.apis.rpc"`
	// ChainAPIGRPCs is a comma seperated list of chain rpc address for various chains, used at output of chain.apis.grpc
	ChainAPIGRPCs string `name:"chain-api-grpcs" json:"chain_api_grpcs" env:"CHAIN_API_GRPCS" usage:"Comma seperated list of chain grpc address for various chains, used at output of chain.apis.grpc"`
	// ChainAPIRESTs is a comma seperated list of chain rpc address for various chains, used at output of chain.apis.rest
	ChainAPIRESTs string `name:"chain-api-rests" json:"chain_api_rests" env:"CHAIN_API_RESTS" usage:"Comma seperated list of chain rest address for various chains, used at output of chain.apis.rest"`
	// ChainClientExposers is a comma seperated list of chain exposer endpoints for various chains
	ChainClientExposers string `name:"chain-client-exposer" json:"chain_client_exposers" env:"CHAIN_CLIENT_EXPOSERS" usage:"Comma seperated list of chain exposer address"`
	// Verbose switches on debug logging
	Verbose bool `name:"verbose" json:"verbose" usage:"switch on debug / verbose logging"`
	// OnlyFatalLog set log level as fatal to ignore logs
	OnlyFatalLog bool `name:"only-fatal-log" json:"only-fatal-log" usage:"used while running test"`
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
			flags = append(flags, cli.BoolTFlag{
				Name:   optName,
				Usage:  msg,
				EnvVar: envName,
			})
		case reflect.String:
			defaultValue := reflect.ValueOf(defaults).Elem().FieldByName(field.Name).String()
			flags = append(flags, cli.StringFlag{
				Name:   optName,
				Usage:  usage,
				EnvVar: envName,
				Value:  defaultValue,
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
	// c.Encoding = "console"

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
