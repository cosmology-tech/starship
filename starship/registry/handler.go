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

	pb "github.com/hyperweb-io/starship/registry/registry"
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

// verifyChainIDs will check the config chain ids are present as directories in
// the chain registry directory
func verifyChainIDs(config *Config) error {
	chainIDs := strings.Split(config.ChainClientIDs, ",")

	files, err := os.ReadDir(config.ChainRegistry)
	if err != nil {
		return err
	}

	var notChainIDs []string
	for _, chainID := range chainIDs {
		var found bool
		for _, file := range files {
			if file.Name() == chainID {
				found = true
				break
			}
		}

		if !found {
			notChainIDs = append(notChainIDs, chainID)
		}
	}

	if len(notChainIDs) > 0 {
		return fmt.Errorf("chains %s not found in chain registry directory", strings.Join(notChainIDs, ","))
	}

	return nil
}

func (a *AppServer) ListChains(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseChains, error) {
	chainIDs := strings.Split(a.config.ChainClientIDs, ",")

	var chains []*pb.ChainRegistry
	for _, chainID := range chainIDs {
		resp, err := a.GetChain(ctx, &pb.RequestChain{Chain: chainID})
		if err != nil {
			return nil, err
		}

		chains = append(chains, resp)
	}

	return &pb.ResponseChains{Chains: chains}, nil
}

func (a *AppServer) ListChainIDs(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseChainIDs, error) {
	chainIDs := strings.Split(a.config.ChainClientIDs, ",")

	return &pb.ResponseChainIDs{ChainIds: chainIDs}, nil
}

// GetChain handles the incoming request for a single chain given the chain id
// Note, we use chain-id instead of chain type, since it is expected, that there
// can be multiple chains of same type by unique chain ids
func (a *AppServer) GetChain(ctx context.Context, requestChain *pb.RequestChain) (*pb.ChainRegistry, error) {
	chainID := requestChain.Chain

	filename := filepath.Join(a.config.ChainRegistry, chainID, "chain.json")
	chain := &pb.ChainRegistry{}

	err := readJSONToProto(filename, chain)
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("unable to read file %s, err: %s", filename, err)
	}

	client, err := a.chainClients.GetChainClient(requestChain.Chain)
	if err != nil {
		return nil, err
	}

	// Fetch and overide peers
	peers, err := a.getChainPeers(ctx, client)
	if err != nil {
		return nil, err
	}
	chain.Peers = peers

	// Fetch and overide apis
	apis, err := a.getChainAPIs(ctx, client)
	if err != nil {
		return nil, err
	}
	chain.Apis = apis

	return chain, nil
}

func (a *AppServer) getChainPeers(ctx context.Context, client *ChainClient) (*pb.Peers, error) {
	seedPeers, err := client.GetChainSeed(ctx)
	if err != nil {
		return nil, err
	}

	persistentPeers, err := client.GetChainPeers(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.Peers{Seeds: seedPeers, PersistentPeers: persistentPeers}, nil
}

// ListChainPeers fetches all the peers for the chain
func (a *AppServer) ListChainPeers(ctx context.Context, requestChain *pb.RequestChain) (*pb.Peers, error) {
	client, err := a.chainClients.GetChainClient(requestChain.Chain)
	if err != nil {
		return nil, err
	}

	return a.getChainPeers(ctx, client)
}

func (a *AppServer) getChainAPIs(ctx context.Context, client *ChainClient) (*pb.APIs, error) {
	chainID := client.ChainID()

	chainIDs := strings.Split(a.config.ChainClientIDs, ",")

	index := 0
	found := false
	for index = range chainIDs {
		if chainIDs[index] == chainID {
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("not found: chain id %s not in configs", chainID)
	}

	moniker, err := client.GetNodeMoniker(ctx)
	if err != nil {
		return nil, err
	}

	apiFactory := func(addrs string) []*pb.APIs_API {
		if addrs == "" {
			return nil
		}
		addrsList := strings.Split(addrs, ",")
		if index > len(addrsList) {
			return nil
		}
		return []*pb.APIs_API{
			{
				Address:  addrsList[index],
				Provider: moniker,
			},
		}
	}

	apis := &pb.APIs{
		Rpc:  apiFactory(a.config.ChainAPIRPCs),
		Grpc: apiFactory(a.config.ChainAPIGRPCs),
		Rest: apiFactory(a.config.ChainAPIRESTs),
	}

	return apis, nil
}

// ListChainAPIs fetches all the apis
func (a *AppServer) ListChainAPIs(ctx context.Context, requestChain *pb.RequestChain) (*pb.APIs, error) {
	chainID := requestChain.Chain
	client, err := a.chainClients.GetChainClient(chainID)
	if err != nil {
		return nil, err
	}

	return a.getChainAPIs(ctx, client)
}

// GetChainKeys fetches all keys for the chain
func (a *AppServer) GetChainKeys(ctx context.Context, requestChain *pb.RequestChain) (*pb.Keys, error) {
	client, err := a.chainClients.GetChainClient(requestChain.Chain)
	if err != nil {
		return nil, err
	}

	return client.GetChainKeys(ctx)
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

		resData = append(resData, infos.ToProto()...)
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

	return &pb.ResponseListIBC{Data: infos.ToProto()}, nil
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

	infoProtos := infos.ToProto()
	for _, info := range infoProtos {
		if info.Chain_2.ChainName == client.GetChainNameFromChainID(requestIBCInfo.Chain_2) {
			return info, nil
		}
	}

	return nil, fmt.Errorf("not found: no ibc connection found between %s and %s", requestIBCInfo.Chain_1, requestIBCInfo.Chain_2)
}
