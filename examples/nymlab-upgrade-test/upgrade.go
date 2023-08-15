package starship

import (
	"context"
	"time"

	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func (s *TestSuite) Upgrade() {
	cheqd := s.GetChainClient("cheqd-testnet-6")

	currHeight, err := cheqd.GetHeight()
	s.Require().NoError(err)

	upgradeName := "v2"
	upgradeHeight := currHeight + 50

	s.T().Logf("submitting v2 upgrade proposal, upgrade height: %d, current height: %d", upgradeHeight, currHeight)
	content := upgradetypes.NewSoftwareUpgradeProposal(
		"cheqd v2 upgrade test",
		"cheqd v2 upgrade test",
		upgradetypes.Plan{
			Name:   upgradeName,
			Height: upgradeHeight,
			Info:   "",
		},
	)
	proposalID := s.SubmitAndVoteProposal(cheqd, content, "upgrade to v8")
	s.T().Logf("proposal submitted: %d", proposalID)

	// timeout_commit is set to 800ms
	blockTime := 800 * time.Millisecond
	expectedTimeToUpgradeHeight := time.Duration(upgradeHeight-currHeight-5) * blockTime // keeping margin for 5 blocks
	// sleeping here because WaitForHeight hits status rest api every second to check height
	// and gets this error after many repetitive calls
	s.T().Logf("Wating for %f seconds", expectedTimeToUpgradeHeight.Seconds())
	time.Sleep(expectedTimeToUpgradeHeight)

	s.T().Log("waiting for upgrade height")
	s.WaitForHeight(cheqd, upgradeHeight)

	s.T().Log("checking proposal status")
	res, err := govv1beta1.
		NewQueryClient(cheqd.Client).
		Proposal(context.Background(), &govv1beta1.QueryProposalRequest{ProposalId: proposalID})
	s.Require().NoError(err)
	s.Require().Equal(govv1beta1.StatusPassed, res.Proposal.Status, "upgrade proposal did not pass before upgrade height: %d", upgradeHeight)

	s.T().Log("verifying upgrade happened")
	planRes, err := upgradetypes.
		NewQueryClient(cheqd.Client).
		AppliedPlan(context.Background(), &upgradetypes.QueryAppliedPlanRequest{Name: upgradeName})
	s.Require().NoError(err)
	s.Require().Equal(upgradeHeight, planRes.Height)
	s.T().Log("upgrade successful")
}
