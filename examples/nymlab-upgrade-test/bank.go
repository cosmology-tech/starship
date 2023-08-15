package starship

import (
	"context"
	"fmt"

	starship "github.com/cosmology-tech/starship/clients/go/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (s *TestSuite) RunTokenTransferTests() {
	cheqd := s.GetChainClient("cheqd-testnet-6")
	ucheq := cheqd.MustGetChainDenom()

	address, err := cheqd.CreateRandWallet("test-transfer")
	s.Require().NoError(err)

	amt := 2345000
	balBefore := s.GetBalance(cheqd, address, ucheq)
	s.T().Logf("transfering %d%s to addr: %s", amt, ucheq, address)
	s.TransferTokens(cheqd, address, amt, ucheq)
	balAfter := s.GetBalance(cheqd, address, ucheq)
	s.T().Log("verifying balance after transfer")
	s.Require().Equal(balBefore.AddAmount(sdk.NewInt(int64(amt))), balAfter)
}

func (s *TestSuite) GetBalance(chain *starship.ChainClient, addr string, denom string) sdk.Coin {
	if denom == "" {
		return sdk.Coin{Amount: sdk.NewInt(0)}
	}
	res, err := banktypes.
		NewQueryClient(chain.Client).
		Balance(context.Background(), &banktypes.QueryBalanceRequest{
			Address: addr,
			Denom:   denom,
		})
	s.Require().NoError(err)
	return *res.GetBalance()
}

func (s *TestSuite) TransferTokens(chain *starship.ChainClient, addr string, amount int, denom string) {
	coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("%d%s", amount, denom))
	s.Require().NoError(err)

	// Build transaction message
	msg := &banktypes.MsgSend{
		FromAddress: chain.Address,
		ToAddress:   addr,
		Amount:      sdk.Coins{coin},
	}
	s.SendMsgAndWait(chain, msg, "Transfer tokens for e2e tests")
}
