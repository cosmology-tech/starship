package main

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/cosmology-tech/starship/faucet/faucet"
)

func (a *AppServer) Status(ctx context.Context, _ *emptypb.Empty) (*pb.State, error) {
	return nil, ErrNotImplemented
}

func (a *AppServer) Credit(ctx context.Context, requestCredit *pb.RequestCredit) (*pb.ResponseCredit, error) {
	return nil, ErrNotImplemented
}
