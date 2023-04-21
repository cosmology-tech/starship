package upgrade_test

import (
	lens "github.com/strangelove-ventures/lens/client"
)

type ChainClient struct {
	logger      *zap.Logger
	config 	*Config

	chainConfig *lens.ChainClientConfig
	client *lens.ChainClient
}

func NewChainClient(logger *zap.Logger, config *Config) (*ChainClient, error) {
	ccc := &lens.ChainClientConfig{
		ChainID:        config.ChainID,
		RPCAddr:        config.RPCAddr,
	}
}
