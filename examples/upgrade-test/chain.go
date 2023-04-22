package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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
	chainConfig *lens.ChainClientConfig
	client      *lens.ChainClient
}

func NewChainClient(logger *zap.Logger, config *Config, chainID string) (*ChainClient, error) {
	cc := config.GetChain(chainID)

	ccc := &lens.ChainClientConfig{
		ChainID:        chainID,
		RPCAddr:        cc.GetRPCAddr(),
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		SignModeStr:    "direct",
	}

	client, err := lens.NewChainClient(logger, ccc, os.Getenv("HOME"), os.Stdin, os.Stdout)
	if err != nil {
		return nil, err
	}

	chainClient := &ChainClient{
		logger:      logger,
		config:      config,
		chainConfig: ccc,
		client:      client,
	}

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
	return c.chainConfig.ChainID
}

func (c *ChainClient) Initialize() error {
	keyName := "genesis"
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

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var keys *pb.Keys
	err = json.Unmarshal(respData, keys)
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
	walletAddr, err := c.client.RestoreKey(keyName, mnemonic, 118)
	if err != nil {
		return "", err
	}

	return walletAddr, nil
}
