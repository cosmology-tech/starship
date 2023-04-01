package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Anmol1696/starship/registry/registry"
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
		if strings.HasPrefix(f.Name(), "_") || !f.IsDir() {
			continue
		}

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
		if strings.HasPrefix(f.Name(), "_") || !f.IsDir() {
			continue
		}

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

// ListChainPeers fetches all the peers for the chain
func (a *AppServer) ListChainPeers(ctx context.Context, requestChain *pb.RequestChain) (*pb.Peers, error) {
	client, err := a.chainClients.GetChainClient(requestChain.Chain)
	if err != nil {
		return nil, err
	}

	return nil, nil
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

// ListIBC will return all the current IBC connections
func (a *AppServer) ListIBC(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseListIBC, error) {
	var resData []*pb.IBCData
	for _, client := range a.chainClients {
		infos, err := client.GetChainInfo()
		if err != nil {
			return nil, err
		}

		for _, info := range infos {
			resData = append(resData, info.ToProto())
		}
	}

	return &pb.ResponseListIBC{Data: resData}, nil
}

// ListChainIBC will return all the current IBC connections
func (a *AppServer) ListChainIBC(ctx context.Context, requestChain *pb.RequestChain) (*pb.ResponseListIBC, error) {
	client, err := a.chainClients.GetChainClient(requestChain.Chain)
	if err != nil {
		return nil, err
	}

	infos, err := client.GetChainInfo()
	if err != nil {
		return nil, err
	}

	var resData []*pb.IBCData
	for _, info := range infos {
		resData = append(resData, info.ToProto())
	}

	return &pb.ResponseListIBC{Data: resData}, nil
}

func (a *AppServer) GetIBCInfo(ctx context.Context, requestIBCInfo *pb.RequestIBCInfo) (*pb.IBCData, error) {
	client, err := a.chainClients.GetChainClient(requestIBCInfo.Chain_1)
	if err != nil {
		return nil, err
	}

	infos, err := client.GetChainInfo()
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		if info.Counterparty.ChainId == requestIBCInfo.Chain_2 {
			return info.ToProto(), nil
		}
	}

	return nil, fmt.Errorf("not found: no ibc connection found between %s and %s", requestIBCInfo.Chain_1, requestIBCInfo.Chain_2)
}
