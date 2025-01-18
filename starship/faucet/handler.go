package main

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hyperweb-io/starship/faucet/faucet"
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

func (a *AppServer) Credit(ctx context.Context, requestCredit *pb.RequestCredit) (*pb.ResponseCredit, error) {
	err := a.distributor.SendTokens(requestCredit.GetAddress(), requestCredit.GetDenom())
	if err != nil {
		return nil, err
	}

	return &pb.ResponseCredit{Status: "ok"}, nil
}
