package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/cosmology-tech/starship/faucet/faucet"
)

func (a *AppServer) Status(ctx context.Context, _ *emptypb.Empty) (*pb.State, error) {
	accountBalances, err := a.distributor.Status()
	if err != nil {
		return nil, err
	}

	holder := &pb.AddressBalance{}
	distributors := []*pb.AddressBalance{}
	for _, accountBalance := range accountBalances {
		if accountBalance.Account.Address == a.distributor.Holder.Address {
			holder = accountBalance.ToProto()
			continue
		}
		distributors = append(distributors, accountBalance.ToProto())
	}

	state := &pb.State{
		Status:          "ok",
		NodeUrl:         a.config.ChainRPCEndpoint,
		ChainId:         a.config.ChainId,
		ChainTokens:     a.distributor.CreditCoins.GetDenoms(),
		AvailableTokens: a.distributor.CreditCoins.GetDenoms(),
		Holder:          holder,
		Distributors:    distributors,
	}

	return state, nil
}

func (a *AppServer) getBalance(address, denom string) (*big.Int, error) {
  account := &Account{config: a.config, logger: a.logger, Address: address}
  coin, err := account.GetBalanceByDenom(denom)
  if err != nil {
    // Log the error, but don't return it
    a.logger.Debug("Error getting balance, assuming new account", zap.Error(err))
    return new(big.Int), nil // Return 0 balance
  }
  balance, ok := new(big.Int).SetString(coin.Amount, 10)
  if !ok {
    return nil, fmt.Errorf("failed to parse balance")
  }
  return balance, nil
}

func (a *AppServer) Credit(ctx context.Context, requestCredit *pb.RequestCredit) (*pb.ResponseCredit, error) {
  // Get initial balance before sending tokens
  initialBalance, err := a.getBalance(requestCredit.GetAddress(), requestCredit.GetDenom())
  if err != nil {
    return nil, fmt.Errorf("failed to get initial balance: %v", err)
  }

  err = a.distributor.SendTokens(requestCredit.GetAddress(), requestCredit.GetDenom())
  if err != nil {
    return nil, err
  }

  // Check balance after transfer
  confirmed, err := a.confirmBalanceUpdate(requestCredit.GetAddress(), requestCredit.GetDenom(), initialBalance)
  if err != nil {
    return &pb.ResponseCredit{Status: fmt.Sprintf("error: %v", err)}, err
  }
  if !confirmed {
    return &pb.ResponseCredit{Status: "error: failed to confirm balance update (timeout)"}, nil
  }

  return &pb.ResponseCredit{Status: "ok"}, nil
}

func (a *AppServer) confirmBalanceUpdate(address, denom string, initialBalance *big.Int) (bool, error) {
  expectedIncrease, ok := new(big.Int).SetString(a.distributor.CreditCoins.GetDenomAmount(denom), 10)
  if !ok {
    return false, fmt.Errorf("failed to parse expected amount")
  }

  expectedFinalBalance := new(big.Int).Add(initialBalance, expectedIncrease)

  for i := 0; i < 3; i++ { // Try 3 times with 5-second intervals
    currentBalance, err := a.getBalance(address, denom)
    if err != nil {
      return false, err
    }
    if currentBalance.Cmp(expectedFinalBalance) >= 0 {
      return true, nil
    }
    if i < 2 {
      time.Sleep(5 * time.Second)
    }
  }
  return false, nil
}
