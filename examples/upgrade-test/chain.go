package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang/protobuf/jsonpb"
	lens "github.com/strangelove-ventures/lens/client"
	"go.uber.org/zap"

	pb "github.com/cosmology-tech/starship/registry/registry"
)

type ChainClients []*ChainClient

func NewChainClients(logger *zap.Logger, config *Config) (ChainClients, error) {
	var clients []*ChainClient
	for _, chain := range config.Chains {
		client, err := NewChainClient(logger, config, chain.Name)
		if err != nil {
			logger.Error("unable to create client for chain",
				zap.String("chain_id", chain.Name),
				zap.Error(err),
			)
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetChainClient returns a chain client pointer for the given chain id
func (cc ChainClients) GetChainClient(chainID string) (*ChainClient, error) {
	for _, client := range cc {
		if client.ChainID() == chainID {
			return client, nil
		}
	}

	return nil, fmt.Errorf("not found: client chain id %s", chainID)
}

type ChainClient struct {
	logger *zap.Logger
	config *Config

	address     string
	chainID     string
	chainConfig *lens.ChainClientConfig
	client      *lens.ChainClient
}

func NewChainClient(logger *zap.Logger, config *Config, chainID string) (*ChainClient, error) {
	cc := config.GetChain(chainID)

	chainClient := &ChainClient{
		logger:  logger,
		config:  config,
		chainID: chainID,
	}

	// fetch chain registry from the local registry
	registry, err := chainClient.GetChainRegistry()
	if err != nil {
		return nil, err
	}

	ccc := &lens.ChainClientConfig{
		ChainID:        chainID,
		RPCAddr:        cc.GetRPCAddr(),
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		SignModeStr:    "direct",
		AccountPrefix:  *registry.Bech32Prefix,
		GasAdjustment:  1.5,
		GasPrices:      fmt.Sprintf("%f%s", registry.Fees.FeeTokens[0].HighGasPrice, registry.Fees.FeeTokens[0].Denom),
		MinGasAmount:   0,
		Slip44:         int(registry.Slip44),
	}

	client, err := lens.NewChainClient(logger, ccc, os.Getenv("HOME"), os.Stdin, os.Stdout)
	if err != nil {
		return nil, err
	}

	chainClient.chainConfig = ccc
	chainClient.client = client

	err = chainClient.Initialize()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func (c *ChainClient) GetRPCAddr() string {
	return c.config.GetChain(c.ChainID()).GetRPCAddr()
}

func (c *ChainClient) ChainID() string {
	return c.chainID
}

func (c *ChainClient) Initialize() error {
	keyName := fmt.Sprintf("genesis-%s", c.ChainID())
	mnemonic, err := c.GetGenesisMnemonic()
	if err != nil {
		return err
	}

	wallet, err := c.GetWallet(keyName, mnemonic)
	if err != nil {
		return err
	}

	c.address = wallet
	c.chainConfig.Key = keyName

	return nil
}

// GetChainKeys fetches keys from the chain registry at `/chains/{chain-id}/keys` endpoint
func (c *ChainClient) GetChainKeys() (*pb.Keys, error) {
	url := fmt.Sprintf("%s/chains/%s/keys", c.config.Registry.GetRESTAddr(), c.ChainID())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	keys := &pb.Keys{}
	err = jsonpb.Unmarshal(resp.Body, keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

// GetGenesisMnemonic fetches the mnemonic from GetChainKeys and returns the first mnemonic in genesis list
func (c *ChainClient) GetGenesisMnemonic() (string, error) {
	keys, err := c.GetChainKeys()
	if err != nil {
		return "", err
	}

	return keys.Genesis[0].Mnemonic, nil
}

func (c *ChainClient) GetWallet(keyName, mnemonic string) (string, error) {
	// delete key if already exists
	//_, err := c.client.DeleteKey(keyName)

	walletAddr, err := c.client.RestoreKey(keyName, mnemonic, 118)
	if err != nil {
		return "", err
	}

	return walletAddr, nil
}

// GetChainRegistry fetches the chain registry from the registry at `/chains/{chain-id}` endpoint
func (c *ChainClient) GetChainRegistry() (*pb.ChainRegistry, error) {
	url := fmt.Sprintf("%s/chains/%s", c.config.Registry.GetRESTAddr(), c.ChainID())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	chainRegistry := &pb.ChainRegistry{}
	err = jsonpb.Unmarshal(resp.Body, chainRegistry)
	if err != nil {
		return nil, err
	}

	// verify chain id from chain registry and config
	if chainRegistry.ChainId != c.ChainID() {
		return nil, fmt.Errorf("chain id mismatch: %s != %s", chainRegistry.ChainId, c.ChainID())
	}

	return chainRegistry, nil
}