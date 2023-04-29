package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/golang/protobuf/jsonpb"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	expb "github.com/cosmology-tech/starship/exposer/exposer"
)

var configFile = "./config.yaml"

type TestSuite struct {
	suite.Suite

	config       *Config
	chainClients ChainClients
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.T().Log("setting up e2e integration test suite...")

	// read config file from yaml
	yamlFile, err := os.ReadFile(configFile)
	s.Require().NoError(err)
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	s.Require().NoError(err)
	s.config = config

	// create chain clients
	chainClients, err := NewChainClients(zap.L(), config)
	s.Require().NoError(err)
	s.chainClients = chainClients
}

func (s *TestSuite) MakeRequest(req *http.Request, expCode int) io.Reader {
	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err, "trying to make request", zap.Any("request", req))

	s.Require().Equal(expCode, resp.StatusCode, "response code did not match")

	return resp.Body
}

// WaitForTx will wait for the tx to complete.
func (s *TestSuite) WaitForTx(chain *ChainClient, txHex string) {
	tx, err := chain.client.QueryTx(context.Background(), txHex, false)
	s.Require().NoError(err)
	s.Assert().NotNil(tx)
	time.Sleep(5 * time.Second)
	tx, err = chain.client.QueryTx(context.Background(), txHex, false)
	s.Require().NoError(err)
	s.Assert().NotNil(tx)
}

// IBCTransferTokens will transfer chain native token from chain1 to chain2 at given address
func (s *TestSuite) IBCTransferTokens(chain1, chain2 *ChainClient, chain2Addr string, amount int) {
	channel, err := chain1.GetIBCChannel(chain2.ChainID())
	s.Require().NoError(err)

	denom, err := chain1.GetChainDenom()
	s.Require().NoError(err)

	coin := sdk.Coin{Denom: denom, Amount: sdk.NewInt(int64(amount))}
	req := &transfertypes.MsgTransfer{
		SourcePort:       channel.Chain_2.PortId,
		SourceChannel:    channel.Chain_2.ChannelId,
		Token:            coin,
		Sender:           chain1.address,
		Receiver:         chain2Addr,
		TimeoutHeight:    clienttypes.NewHeight(12300, 45600),
		TimeoutTimestamp: 0,
		Memo:             fmt.Sprintf("testsetup: transfer token from %s to %s", chain1.ChainID(), chain2.ChainID()),
	}

	res, err := chain1.client.SendMsg(context.Background(), req, "")
	s.Require().NoError(err)
	if err != nil {
		s.Require().Nil(res, "msg failed", zap.Any("error", err), zap.Any("code", res.Code), zap.Any("logs", res.Logs))
	}
}

func (s *TestSuite) TestChains_Status() {
	s.T().Log("runing test for /status endpoint for each chain")

	for _, chainClient := range s.chainClients {
		url := fmt.Sprintf("%s/status", chainClient.GetRPCAddr())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		s.Require().NoError(err)

		body := s.MakeRequest(req, 200)
		resp := &expb.Status{}
		err = jsonpb.Unmarshal(body, resp)
		s.Assert().NoError(err)

		// assert chain id
		s.Assert().Equal(chainClient.ChainID(), resp.Result.NodeInfo.Network)
	}
}

func (s *TestSuite) TestChain_TokenTransfer() {
	chain1, err := s.chainClients.GetChainClient("core-1")
	s.Require().NoError(err)

	keyName := "test-transfer"
	address, err := chain1.CreateRandWallet(keyName)
	s.Require().NoError(err)

	denom, err := chain1.GetChainDenom()
	s.Require().NoError(err)
	coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("2345000%s", denom))
	s.Require().NoError(err)
	// Build transaction message
	req := &banktypes.MsgSend{
		FromAddress: chain1.address,
		ToAddress:   address,
		Amount:      sdk.Coins{coin},
	}
	res, err := chain1.client.SendMsg(context.Background(), req, "")
	s.Require().NoError(err)

	s.WaitForTx(chain1, res.TxHash)

	s.T().Log("response recived", zap.Any("response", res))

	// Verify the address recived the token
	balance, err := chain1.client.QueryBalanceWithDenomTraces(context.Background(), sdk.MustAccAddressFromBech32(address), nil)
	s.Require().NoError(err)

	// Assert correct transfers
	s.Assert().Len(balance, 1)
	s.Assert().Equal(balance.Denoms(), []string{denom})
	s.Assert().Equal(balance[0].Amount, sdk.NewInt(2345000))
}

func (s *TestSuite) TestIBCTransfer() {
	chain1, err := s.chainClients.GetChainClient("core-1")
	s.Require().NoError(err)
	chain2, err := s.chainClients.GetChainClient("gaia-4")
	s.Require().NoError(err)

	keyName := "test-ibc-transfer"
	address, err := chain1.CreateRandWallet(keyName)
	s.Require().NoError(err)

	// Tranfer atom to persistence chain
	s.IBCTransferTokens(chain2, chain1, address, 12345000)

	// Verify the address recived the token
	balance, err := chain1.client.QueryBalanceWithDenomTraces(context.Background(), sdk.AccAddress(address), nil)
	s.Require().NoError(err)

	chain2Denom, err := chain2.GetChainDenom()
	s.Require().NoError(err)

	// Assert correct transfers
	s.Assert().Len(balance, 1)
	s.Assert().Equal(balance.Denoms(), []string{chain2Denom})
	s.Assert().Equal(balance[0].Amount, sdk.NewInt(12345000))
}
