package main

import (
	"context"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// TestChainUpgrade is the overall test for performing a software upgrade
func (s *TestSuite) TestChainUpgrade() {
	if testing.Short() {
		s.T().Skip("Skipping chain upgrade tests for short test")
	}

	chain, err := s.chainClients.GetChainClient("core-1")
	s.Require().NoError(err)
	keyName := "upgrade-test-address"
	address, err := chain.CreateRandWallet(keyName)
	s.Require().NoError(err)

	// Perform pre-upgrade steps
	// Transfer tokens to address
	denom := chain.MustGetChainDenom()
	s.TransferTokens(chain, address, 12312300, denom)

	// Perform chain upgrade
	// fetch all the information needed for the upgrade
	cc := chain.config.GetChain("core-1")
	version := cc.Upgrade.Upgrades[0].Name
	curHeight, err := chain.GetHeight()
	s.Require().NoError(err)
	upgradeHeight := curHeight + 100

	// Create software upgrade proposal
	plan := upgradetypes.Plan{
		Name:   version,
		Height: upgradeHeight,
		Info:   "",
	}
	content := upgradetypes.NewSoftwareUpgradeProposal(
		"Software upgrade",
		"software upgrade",
		plan,
	)

	msg := &gov.MsgSubmitProposal{
		InitialDeposit: sdk.NewCoins(sdk.NewCoin("uxprt", sdk.NewInt(10000000))),
		Proposer:       chain.address,
	}
	err = msg.SetContent(content)
	s.Require().NoError(err)
	// Submit proposal
	res, err := chain.SendMsg(context.Background(), msg, "Software upgrade proposal")
	s.Require().NoError(err)
	s.WaitForTx(chain, res.TxHash)

	// Fetch the proposal id
	proposalID := 0
	for _, event := range res.Logs[0].Events {
		if event.Type == "submit_proposal" {
			for _, attr := range event.Attributes {
				if attr.Key == "proposal_id" {
					value, err := strconv.Atoi(attr.Value)
					s.Require().NoError(err)
					proposalID = value
				}
			}
		}
	}
	s.Require().NotZerof(proposalID, "proposal id not found in tx events")

	// Vote on the proposal to pass it
	vote := &gov.MsgVote{ProposalId: uint64(proposalID), Voter: chain.address, Option: gov.OptionYes}
	res, err = chain.SendMsg(context.Background(), vote, "Votes for software upgrade proposal")
	s.Require().NoError(err)
	s.WaitForTx(chain, res.TxHash)

	// Wait for upgrade height

	// Query proposal to see if it is in voting period
	req := &gov.QueryProposalRequest{ProposalId: uint64(proposalID)}
	queryclient := gov.NewQueryClient(chain.client)
	propRes, err := queryclient.Proposal(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(propRes)
	s.Require().Equal(propRes.Proposal.ProposalId, uint64(proposalID))
	s.Require().Equal(propRes.Proposal.Status, gov.StatusVotingPeriod)

	// Wait for upgrade height and fetch propsal again
	s.WaitForHeight(chain, upgradeHeight+1)

	propRes, err = queryclient.Proposal(context.Background(), req)
	s.Require().NoError(err)
	s.Require().NotNil(propRes)
	s.Require().Equal(propRes.Proposal.Status, gov.StatusPassed)

	// Verify upgrade happened
	upgradeClient := upgradetypes.NewQueryClient(chain.client)
	planRes, err := upgradeClient.AppliedPlan(context.Background(), &upgradetypes.QueryAppliedPlanRequest{Name: version})
	s.Require().NoError(err)
	s.Require().Equal(planRes.Height, upgradeHeight)

	// Perform post upgrade tests/checks
	// balance of address
	balance, err := chain.client.QueryBalanceWithDenomTraces(context.Background(), sdk.MustAccAddressFromBech32(address), nil)
	s.Require().NoError(err)
	// Assert correct transfers
	s.Assert().Len(balance, 1)
	s.Assert().Equal(balance.Denoms(), []string{denom})
	s.Assert().Equal(balance[0].Amount, sdk.NewInt(12312300))

	// transfer some more tokens to address
	s.TransferTokens(chain, address, 12312300, denom)

	// Check balance again
	balance, err = chain.client.QueryBalanceWithDenomTraces(context.Background(), sdk.MustAccAddressFromBech32(address), nil)
	s.Require().NoError(err)
	// Assert correct transfers
	s.Assert().Len(balance, 1)
	s.Assert().Equal(balance.Denoms(), []string{denom})
	s.Assert().Equal(balance[0].Amount, sdk.NewInt(24624600))
}
