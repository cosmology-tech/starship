package starship

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/suite"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	starship "github.com/cosmology-tech/starship/clients/go/client"
)

var configFile = "./config.yaml"

type TestSuite struct {
	suite.Suite

	config       *starship.Config
	chainClients starship.ChainClients
}

func (s *TestSuite) SetupTest() {
	s.T().Log("setting up e2e integration test suite...")

	// read config file from yaml
	yamlFile, err := os.ReadFile(configFile)
	s.Require().NoError(err)
	config := &starship.Config{}
	err = yaml.Unmarshal(yamlFile, config)
	s.Require().NoError(err)
	s.config = config

	// create chain clients
	chainClients, err := starship.NewChainClients(zap.L(), config)
	s.Require().NoError(err)
	s.chainClients = chainClients
}

func (s *TestSuite) GetChainClient(chainID string) *starship.ChainClient {
	chain, err := s.chainClients.GetChainClient(chainID)
	s.Require().NoError(err)
	return chain
}

func (s *TestSuite) SendMsgAndWait(chain *starship.ChainClient, msg sdk.Msg, memo string) *sdk.TxResponse {
	res, err := chain.Client.SendMsg(context.Background(), msg, memo)
	s.Require().NoError(err)
	s.WaitForTx(chain, res.TxHash)
	return res
}

// WaitForTx will wait for the tx to complete, fail if not able to find tx
func (s *TestSuite) WaitForTx(chain *starship.ChainClient, txHex string) {
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

func (s *TestSuite) WaitForHeight(chain *starship.ChainClient, height int64) {
	s.T().Logf("waiting for height: %d", height)
	s.Require().Eventuallyf(
		func() bool {
			curHeight, err := chain.GetHeight()
			// retry if error is of EOF
			// sometimes this happens with error
			// post failed: Post "http://localhost:26657": EOF
			if errors.Is(err, io.EOF) {
				time.Sleep(time.Second) // add some delay in next call
				return false
			}
			s.Require().NoError(err)
			return curHeight >= height
		},
		300*time.Second,
		time.Second,
		"waited for too long, still height did not reach desired block height",
	)
}

func (s *TestSuite) WaitForNextBlock(chain *starship.ChainClient) {
	currHeight, err := chain.GetHeight()
	s.Require().NoError(err)
	s.WaitForHeight(chain, currHeight+1)
}

func (s *TestSuite) WaitForProposalToPass(chain *starship.ChainClient, proposalID uint64) {
	s.Require().Eventuallyf(
		func() bool {
			res, err := govv1beta1.
				NewQueryClient(chain.Client).
				Proposal(context.Background(), &govv1beta1.QueryProposalRequest{ProposalId: proposalID})
			s.Require().NoError(err)
			return res != nil && res.Proposal.Status == govv1beta1.StatusPassed
		},
		300*time.Second,
		time.Second,
		"waited for too long, proposal is still not passed",
	)
}

func (s *TestSuite) SubmitAndVoteProposal(chain *starship.ChainClient, content govv1beta1.Content, memo string) uint64 {
	denom := chain.MustGetChainDenom()
	msg := &govv1beta1.MsgSubmitProposal{
		InitialDeposit: sdk.NewCoins(sdk.NewCoin(denom, sdk.NewInt(10000000))),
		Proposer:       chain.Address,
	}
	err := msg.SetContent(content)
	s.Require().NoError(err)
	s.T().Logf("submitting proposal: %s", memo)
	res := s.SendMsgAndWait(chain, msg, memo)

	id := s.FindEventAttr(res, "submit_proposal", "proposal_id")
	proposalID, err := strconv.Atoi(id)
	s.Require().NoError(err)

	s.T().Logf("submitting vote on proposal: %d | memo: %s", proposalID, memo)
	vote := &govv1beta1.MsgVote{ProposalId: uint64(proposalID), Voter: chain.Address, Option: govv1beta1.OptionYes}
	s.SendMsgAndWait(chain, vote, fmt.Sprintf("vote: %s", memo))
	return uint64(proposalID)
}

func (s *TestSuite) SubmitAndPassProposal(chain *starship.ChainClient, content govv1beta1.Content, memo string) {
	proposalID := s.SubmitAndVoteProposal(chain, content, memo)
	s.T().Logf("waiting for proposal to pass: %d | memo: %s", proposalID, memo)
	s.WaitForProposalToPass(chain, proposalID)
}

func (s *TestSuite) FindEventAttr(res *sdk.TxResponse, event, attr string) string {
	for _, txEvent := range res.Logs[0].Events {
		if txEvent.Type == event {
			for _, txAttr := range txEvent.Attributes {
				if txAttr.Key == attr {
					return txAttr.Value
				}
			}
		}
	}
	s.FailNow("event attr not found in tx events")
	return ""
}
