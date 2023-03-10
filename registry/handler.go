package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "registry/registry"
)

func readJSONFile(file string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening json file: %s", file)
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result, nil
}

func readJSONToProto(file string, m proto.Message) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("error opening json file: %s", file)
	}

	err = jsonpb.Unmarshal(jsonFile, m)
	if err != nil {
		return err
	}

	return nil
}

func (a *AppServer) ListChains(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseChains, error) {
	files, err := os.ReadDir(a.config.ChainRegistry)
	if err != nil {
		return nil, err
	}

	var chains []*pb.ChainRegistry
	for _, f := range files {
		chain := &pb.ChainRegistry{}
		filename := filepath.Join(a.config.ChainRegistry, f.Name(), "chain.json")

		err := readJSONToProto(filename, chain)
		if err != nil {
			return nil, err
		}

		chains = append(chains, chain)
	}

	return &pb.ResponseChains{Chains: chains}, err
}

func (a *AppServer) ListChainIDs(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseChainIDs, error) {
	files, err := os.ReadDir(a.config.ChainRegistry)
	if err != nil {
		return nil, err
	}

	var chainIDs []string
	for _, f := range files {
		filename := filepath.Join(a.config.ChainRegistry, f.Name(), "chain.json")
		info, err := readJSONFile(filename)
		if err != nil {
			return nil, err
		}
		chainID, ok := info["chain_id"].(string)
		if !ok {
			return nil, fmt.Errorf("unable to get chain id for %s, err: %s", filename, err)
		}
		chainIDs = append(chainIDs, chainID)
	}

	return &pb.ResponseChainIDs{ChainIds: chainIDs}, nil
}

// GetChain handles the incoming request for a single chain given the chain id
// Note, we use chain-id instead of chain type, since it is expected, that there
// can be multiple chains of same type by unique chain ids
func (a *AppServer) GetChain(ctx context.Context, requestChain *pb.RequestChain) (*pb.ResponseChain, error) {
	chainID := requestChain.Chain

	filename := filepath.Join(a.config.ChainRegistry, chainID, "chain.json")
	chain := &pb.ChainRegistry{}

	err := readJSONToProto(filename, chain)
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("unable to read file %s, err: %d", filename, err)
	}

	return &pb.ResponseChain{Chain: chain}, nil
}

func (a *AppServer) GetChainAssets(ctx context.Context, requestChain *pb.RequestChain) (*pb.ResponseChainAssets, error) {
	chainID := requestChain.Chain

	filename := filepath.Join(a.config.ChainRegistry, chainID, "assetlist.json")
	chainAsset := &pb.ResponseChainAssets{}

	err := readJSONToProto(filename, chainAsset)
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("unable to read file %s, err: %d", filename, err)
	}

	return chainAsset, nil
}

func (a *AppServer) GetAllIBC(ctx context.Context) error {
	return ErrNotImplemented
}

func (a *AppServer) GetIBCChainsData(ctx context.Context) error {
	return ErrNotImplemented
}

func (a *AppServer) SetIBCChainsData(ctx context.Context) error {
	return ErrNotImplemented
}

func (a *AppServer) GetIBCChainsChannels(ctx context.Context) error {
	return ErrNotImplemented
}

func (a *AppServer) AddIBCChainChannel(ctx context.Context) error {
	return ErrNotImplemented
}
