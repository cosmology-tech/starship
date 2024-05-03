package main

import (
	"fmt"
	"reflect"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func NewDefaultConfig() *Config {
	return &Config{
		Host:              "0.0.0.0",
		HTTPPort:          "8000",
		GRPCPort:          "9000",
		ChainRESTEndpoint: "http://localhost:1317",
		ChainRPCEndpoint:  "http://localhost:26657",
		ChainBalancesURI:  "/cosmos/bank/v1beta1/balances",
		Concurrency:       1,
		RefillFactor:      8,
		RefillThreshold:   20,
		RefillEpoch:       100,
		Verbose:           true,
	}
}

type Config struct {
	// Host is the interface to bind the HTTP service on
	Host string `name:"host" json:"host" env:"HOST" usage:"host address to listen on"`
	// HTTPPort is the port for the http server
	HTTPPort string `name:"http-port" json:"http_port" env:"HTTP_PORT" usage:"port for http server"`
	// GRPCPort is the port for the grpc server
	GRPCPort string `name:"grpc-port" json:"grpc_port" env:"GRPC_PORT" usage:"port for gRPC server"`
	// ChainHome is the path the home directory for chain on node
	ChainHome string `name:"chain-home" json:"chain_home" env:"CHAIN_HOME" usage:"path to the home of chain node"`
	// ChainBinary is the binary for running the chain nodes
	ChainBinary string `name:"chain-binary" json:"chain_binary" env:"CHAIN_BINARY" usage:"chain binary name of the same node"`
	// ChainId is the chain id for the given chain
	ChainId string `name:"chain-id" json:"chain_id" env:"CHAIN_ID" usage:"chain id of the given chain"`
	// ChainRESTEndpoint is the chain rest endpoint
	ChainRESTEndpoint string `name:"chain-rest-endpoint" json:"chain_rest_endpoint" env:"CHAIN_REST_ENDPOINT" usage:"lcd endpoint of the chain"`
	// ChainRPCEndpoint is the chain rpc endpoint
	ChainRPCEndpoint string `name:"chain-rpc-endpoint" json:"chain_rpc_endpoint" env:"CHAIN_RPC_ENDPOINT" usage:"rpc endpoint of the chain"`
	// ChainBalancesURI is the uri for getting the balance from the rest endpoint
	ChainBalancesURI string `name:"chain-balances-uri" json:"chain_blanaces_uri" env:"CHAIN_BALANCE_URI" usage:"balances endpoint {chain-balances-uri}/{address} should return balances"`
	// Concurrency is the number of distributor address to use for handing requests
	Concurrency int `name:"concurrency" json:"concurrency" env:"CONCURRENCY" usage:"number of distributor address to use for handling requests"`
	// ChainFees is the amount of fees to provide for the txns
	ChainFees string `name:"chain-fees" json:"chain_fees" env:"CHAIN_FEES" usage:"amount of fees for all txns"`
	// RefillFactor is the factor which times credit amount is sent to the distributors
	RefillFactor int `name:"refill-factor" json:"refill_factor" env:"REFILL_FACTOR" usage:"send factor times credit amount on refilling"`
	// RefillThreshold is the factor which times credit amount is the min balance after which refil will be triggered
	RefillThreshold int `name:"refill-threshold" json:"refill_threshold" env:"REFILL_THRESHOLD" usage:"refill when balance gets below factor times credit amount"`
	// RefillEpoch is the time (in secs) to sleep, after which holder will refill distributor address if needed
	RefillEpoch int `name:"refill-epoch" json:"refill_epoch" env:"REFILL_EPOCH" usage:"after every epoch, holder will distrubute tokens to distributor addresses if required"`
	// CreditCoins is comma seperated list of amount and denom of tokens to be transfered
	CreditCoins string `name:"credit-coins" json:"credit_coins" env:"CREDIT_COINS" usage:"comma seperated list of amount and denom of tokens to be transfered. eg: 10000000uosmo,1000000uion"`
	// Mnemonic is the mnemonic of the address used as the faucet address
	Mnemonic string `name:"mnemonic" json:"mnemonic" env:"MNEMONIC" usage:"mnemonic of the address used as the faucet address"`
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
		case reflect.Int:
			defaultValue := reflect.ValueOf(defaults).Elem().FieldByName(field.Name).Int()
			flags = append(flags, cli.Int64Flag{
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
	// iterate the config and grab command line options via reflection
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
			case reflect.Int:
				reflect.ValueOf(config).Elem().FieldByName(field.Name).SetInt(cx.Int64(name))
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
