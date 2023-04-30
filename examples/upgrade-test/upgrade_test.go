package main

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// TestChainUpgrade is the overall test for performing a software upgrade
func (s *TestSuite) TestChainUpgrade() {
	chain, err := s.chainClients.GetChainClient("core-1")
	s.Require().NoError(err)
	keyName := "upgrade-test-address"
	address, err := chain.CreateRandWallet(keyName)
	s.Require().NoError(err)

	// Transfer tokens to address
	denom, err := chain.GetChainDenom()
	s.Require().NoError(err)
	coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("2345000%s", denom))
	s.Require().NoError(err)
	// Build transaction message
	req := &banktypes.MsgSend{
		FromAddress: chain.address,
		ToAddress:   address,
		Amount:      sdk.Coins{coin},
	}
	_, err = chain.client.SendMsg(context.Background(), req, "")
	s.Require().NoError(err)
}
