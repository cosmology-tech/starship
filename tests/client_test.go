package tests

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	starship "github.com/cosmology-tech/starship/clients/go/client"
)

type TestClientSuite struct {
	suite.Suite

	config       *starship.Config
	chainClients starship.ChainClients
	cmdRunner    *starship.CmdRunner
}

func TestE2ETestClientSuite(t *testing.T) {
	suite.Run(t, new(TestClientSuite))
}

func (s *TestClientSuite) SetupTest() {
	s.T().Log("setting up e2e integration test suite for starship go client")

	// read config file from yaml
	configFile := os.Getenv(configEnvKey)
	configFile = strings.Replace(configFile, "tests/", "", -1)
	yamlFile, err := os.ReadFile(configFile)
	s.Require().NoError(err)
	config := &starship.Config{}
	err = yaml.Unmarshal(yamlFile, config)
	s.Require().NoError(err)

	s.config = config

	// create chain clients
	chainClient, err := starship.NewChainClients(zap.L(), config)
	s.Require().NoError(err)
	s.chainClients = chainClient

	// create cmdRunner
	cmdRunner, err := starship.NewCmdRunner(zap.L(), config)
	s.Require().NoError(err)
	s.cmdRunner = cmdRunner
}

// WaitForTx will wait for the tx to complete, fail if not able to find tx
func (s *TestClientSuite) WaitForTx(chain *starship.ChainClient, txHex string) {
	var tx *coretypes.ResultTx
	var err error
	s.Require().Eventuallyf(
		func() bool {
			tx, err = chain.Client.QueryTx(context.Background(), txHex, false)
			if err != nil {
				return false
			}
			if tx.TxResult.Code == 0 {
				return true
			}
			return false
		},
		300*time.Second,
		time.Second,
		"waited for too long, still txn not successfull",
	)
	s.Assert().NotNil(tx)
}

func (s *TestClientSuite) TestChainClient_Status() {
	for _, chainClient := range s.chainClients {
		status, err := chainClient.GetStatus()
		s.Require().NoError(err)
		s.Require().Equal(chainClient.GetChainID(), status.NodeInfo.Network)
	}
}

func (s *TestClientSuite) TestChainClient_TokenTransfer() {
	chain := s.chainClients[0]

	keyName := "test-transfer"
	amount := 1230000
	address, err := chain.CreateRandWallet(keyName)
	s.Require().NoError(err)

	denom, err := chain.GetChainDenom()
	s.Require().NoError(err)

	// transfer token to address from genesis address
	coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("%d%s", amount, denom))
	s.Require().NoError(err)

	req := &banktypes.MsgSend{
		FromAddress: chain.Address,
		ToAddress:   address,
		Amount:      sdk.Coins{coin},
	}
	res, err := chain.Client.SendMsg(context.Background(), req, "transfer tokens e2e test")
	s.Require().NoError(err)
	// wait for tx
	s.WaitForTx(chain, res.TxHash)

	// verifty the address has recived the token
	balance, err := chain.Client.QueryBalanceWithDenomTraces(
		context.Background(),
		sdk.MustAccAddressFromBech32(address),
		nil)
	s.Require().NoError(err)

	// assert correct amounts
	s.Assert().Len(balance, 1)
	s.Assert().Equal(balance.Denoms(), []string{denom})
	s.Assert().Equal(balance[0].Amount, sdk.NewInt(int64(amount)))
}

func (s *TestClientSuite) TestChainClient_PodName() {
	for _, chain := range s.chainClients {
		podName, err := s.cmdRunner.GetPodFromName(chain.ChainID)
		s.Require().NoError(err)

		s.Assert().Equal(fmt.Sprintf("%s-genesis-0", chain.ChainID), podName)
	}
}

// TestChainClient_CreateChannel will run an exec command on the relayer
// to create a new channel between the chains, then verify the channel was created
func (s *TestClientSuite) TestChainClient_CreateChannel() {
	// only run the test if a relayer exists
	if s.config.Relayers == nil {
		s.T().Skipf("skiping create channel test for configs without relayers")
	}

	chain, err := s.chainClients.GetChainClient(s.config.Relayers[0].Chains[0])
	s.Require().NoError(err)

	ibcData, err := chain.GetIBCInfo(s.config.Relayers[0].Chains[1])
	s.Require().NoError(err)

	// create new transfer ports between chains
	cmd := fmt.Sprintf("hermes create channel --a-chain %s --a-connection %s --a-port transfer --b-port transfer --yes", s.config.Relayers[0].Chains[0], ibcData.Chain_1.ConnectionId)
	err = s.cmdRunner.RunExec(s.config.Relayers[0].Name, cmd)
	s.Assert().NoError(err)

	// check new ibc transfer channel has been created
	// there mush be 2 channels
	afterIbcData, err := chain.GetIBCInfo(s.config.Relayers[0].Chains[1])
	s.Assert().NoError(err)

	// verify the 1 more channel has been created compared to previous state
	s.Assert().Equal(len(ibcData.Channels)+1, len(afterIbcData.Channels))
}
