package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
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
		if client.GetChainID() == chainID {
			return client, nil
		}
	}

	return nil, fmt.Errorf("not found: client chain id %s", chainID)
}

type ChainClient struct {
	Logger *zap.Logger
	Config *Config

	Address     string
	ChainID     string
	ChainConfig *lens.ChainClientConfig
	Client      *lens.ChainClient
}

func NewChainClient(logger *zap.Logger, config *Config, chainID string) (*ChainClient, error) {
	cc := config.GetChain(chainID)

	chainClient := &ChainClient{
		Logger:  logger,
		Config:  config,
		ChainID: chainID,
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
		MinGasAmount:   10000,
		Slip44:         int(registry.Slip44),
		Modules:        lens.ModuleBasics,
	}

	client, err := lens.NewChainClient(logger, ccc, os.Getenv("HOME"), os.Stdin, os.Stdout)
	if err != nil {
		return nil, err
	}

	chainClient.ChainConfig = ccc
	chainClient.Client = client

	err = chainClient.Initialize()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func (c *ChainClient) GetRPCAddr() string {
	return c.Config.GetChain(c.GetChainID()).GetRPCAddr()
}

func (c *ChainClient) GetChainID() string {
	return c.ChainID
}

func (c *ChainClient) Initialize() error {
	keyName := fmt.Sprintf("genesis-%s", c.GetChainID())
	mnemonic, err := c.GetGenesisMnemonic()
	if err != nil {
		return err
	}

	wallet, err := c.CreateWallet(keyName, mnemonic)
	if err != nil {
		return err
	}

	c.Address = wallet
	c.ChainConfig.Key = keyName

	return nil
}

// GetChainKeys fetches keys from the chain registry at `/chains/{chain-id}/keys` endpoint
func (c *ChainClient) GetChainKeys() (*pb.Keys, error) {
	url := fmt.Sprintf("%s/chains/%s/keys", c.Config.Registry.GetRESTAddr(), c.GetChainID())
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

func (c *ChainClient) CreateWallet(keyName, mnemonic string) (string, error) {
	// delete key if already exists
	_ = c.Client.DeleteKey(keyName)

	walletAddr, err := c.Client.RestoreKey(keyName, mnemonic, 118)
	if err != nil {
		return "", err
	}

	return walletAddr, nil
}

func (c *ChainClient) CreateRandWallet(keyName string) (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	address, err := c.CreateWallet(keyName, mnemonic)
	if err != nil {
		return "", err
	}

	return address, nil
}

// GetChainRegistry fetches the chain registry from the registry at `/chains/{chain-id}` endpoint
func (c *ChainClient) GetChainRegistry() (*pb.ChainRegistry, error) {
	url := fmt.Sprintf("%s/chains/%s", c.Config.Registry.GetRESTAddr(), c.GetChainID())
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
	if chainRegistry.ChainId != c.GetChainID() {
		return nil, fmt.Errorf("chain id mismatch: %s != %s", chainRegistry.ChainId, c.GetChainID())
	}

	return chainRegistry, nil
}

func (c *ChainClient) GetChainDenom() (string, error) {
	registry, err := c.GetChainRegistry()
	if err != nil {
		return "", err
	}

	return registry.Staking.StakingTokens[0].Denom, nil
}

func (c *ChainClient) MustGetChainDenom() string {
	denom, err := c.GetChainDenom()
	if err != nil {
		panic(err)
	}
	return denom
}

// GetChainAssets fetches the assets from chain registry at `/chains/{chain-id}/assets` endpoint
func (c *ChainClient) GetChainAssets() ([]*pb.ChainAsset, error) {
	url := fmt.Sprintf("%s/chains/%s/assets", c.Config.Registry.GetRESTAddr(), c.GetChainID())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	respAssets := &pb.ResponseChainAssets{}
	err = jsonpb.Unmarshal(resp.Body, respAssets)
	if err != nil {
		return nil, err
	}

	return respAssets.Assets, nil
}

func (c *ChainClient) GetIBCInfo(chain2 string) (*pb.IBCData, error) {
	url := fmt.Sprintf("%s/ibc/%s/%s", c.Config.Registry.GetRESTAddr(), c.GetChainID(), chain2)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	ibcData := &pb.IBCData{}
	err = jsonpb.Unmarshal(resp.Body, ibcData)
	if err != nil {
		return nil, err
	}

	return ibcData, nil
}

func (c *ChainClient) GetIBCChannel(chain2 string) (*pb.ChannelData, error) {
	ibcInfo, err := c.GetIBCInfo(chain2)
	if err != nil {
		return nil, err
	}

	return ibcInfo.Channels[0], nil
}

func (c *ChainClient) GetStatus() (*coretypes.ResultStatus, error) {
	status, err := c.Client.RPCClient.Status(context.Background())
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (c *ChainClient) GetHeight() (int64, error) {
	status, err := c.GetStatus()
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

func (c *ChainClient) CustomSendMsg(ctx context.Context, keyName string, msg sdk.Msg, memo string) (*sdk.TxResponse, error) {
	return c.CustomSendMsgs(ctx, keyName, []sdk.Msg{msg}, memo)
}

// CustomSendMsgs wraps the msgs in a StdTx, signs and sends it. An error is returned if there
// was an issue sending the transaction. A successfully sent, but failed transaction will
// not return an error. If a transaction is successfully sent, the result of the execution
// of that transaction will be logged. A boolean indicating if a transaction was successfully
// sent and executed successfully is returned.
func (c *ChainClient) CustomSendMsgs(ctx context.Context, keyName string, msgs []sdk.Msg, memo string) (*sdk.TxResponse, error) {
	cc := c.Client
	txf, err := cc.PrepareFactory(cc.TxFactory())
	if err != nil {
		return nil, err
	}

	if memo != "" {
		txf = txf.WithMemo(memo)
	}

	// Set the gas amount on the transaction factory
	adjusted := uint64(1000000)
	txf = txf.WithGas(adjusted)

	// Build the transaction builder
	txb, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		return nil, err
	}

	// Attach the signature to the transaction
	// c.LogFailedTx(nil, err, msgs)
	// Force encoding in the chain specific address
	for _, msg := range msgs {
		cc.Codec.Marshaler.MustMarshalJSON(msg)
	}

	err = func() error {
		done := cc.SetSDKContext()
		// ensure that we allways call done, even in case of an error or panic
		defer done()
		if err = tx.Sign(txf, keyName, txb, false); err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		return nil, err
	}

	// Generate the transaction bytes
	txBytes, err := cc.Codec.TxConfig.TxEncoder()(txb.GetTx())
	if err != nil {
		return nil, err
	}

	// Broadcast those bytes
	res, err := cc.BroadcastTx(ctx, txBytes)
	if err != nil {
		return nil, err
	}

	// transaction was executed, log the success or failure using the tx response code
	// NOTE: error is nil, logic should use the returned error to determine if the
	// transaction was successfully executed.
	if res.Code != 0 {
		return res, fmt.Errorf("transaction failed with code: %d", res.Code)
	}

	return res, nil
}

func (c *ChainClient) SendMsg(ctx context.Context, msg sdk.Msg, memo string) (*sdk.TxResponse, error) {
	return c.SendMsgs(ctx, []sdk.Msg{msg}, memo)
}

// SendMsgs wraps the msgs in a StdTx, signs and sends it. An error is returned if there
// was an issue sending the transaction. A successfully sent, but failed transaction will
// not return an error. If a transaction is successfully sent, the result of the execution
// of that transaction will be logged. A boolean indicating if a transaction was successfully
// sent and executed successfully is returned.
func (c *ChainClient) SendMsgs(ctx context.Context, msgs []sdk.Msg, memo string) (*sdk.TxResponse, error) {
	return c.CustomSendMsgs(ctx, c.ChainConfig.Key, msgs, memo)
}
