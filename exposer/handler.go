package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "exposer/exposer"
)

func fetchNodeStatus(url string) (StatusResponse, error) {
	var statusResp StatusResponse

	resp, err := http.Get(url)
	if err != nil {
		return statusResp, fmt.Errorf("unable to fetch status, err: %d", err)
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return statusResp, fmt.Errorf("unable to parse status response, err: %d", err)
	}

	return statusResp, nil
}

func (a *AppServer) readJSONFile(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		a.logger.Error("Error opening file",
			zap.String("file", filePath),
			zap.Error(err))
		return nil, fmt.Errorf("error opening json file: %s", filePath)
	}

	return io.ReadAll(jsonFile)
}

func (a *AppServer) GetNodeID(ctx context.Context, _ *emptypb.Empty) (*pb.ResponseNodeID, error) {
	status, err := fetchNodeStatus(a.config.StatusURL)
	if err != nil {
		return nil, err
	}

	return &pb.ResponseNodeID{NodeId: status.Result.NodeInfo.ID}, nil
}

func (a *AppServer) GetPubKey(ctx context.Context, _ *emptypb.Empty) (*pb.ResponsePubKey, error) {
	status, err := fetchNodeStatus(a.config.StatusURL)
	if err != nil {
		return nil, err
	}

	resPubKey := &pb.ResponsePubKey{
		Type: "/cosmos.crypto.ed25519.PubKey",
		Key:  status.Result.ValidatorInfo.PubKey.Value,
	}

	return resPubKey, nil
}

func (a *AppServer) GetGenesisFile(ctx context.Context, _ *emptypb.Empty) (*pb.GenesisState, error) {
	jsonFile, err := os.Open(a.config.GenesisFile)
	if err != nil {
		return nil, err
	}

	state := &pb.GenesisState{}
	err = jsonpb.Unmarshal(jsonFile, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (a *AppServer) GetKeys(ctx context.Context, _ *emptypb.Empty) (*pb.Keys, error) {
	jsonFile, err := os.Open(a.config.MnemonicFile)
	if err != nil {
		return nil, err
	}

	keys := &pb.Keys{}
	err = jsonpb.Unmarshal(jsonFile, keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (a *AppServer) GetPrivKeysFile(ctx context.Context, _ *emptypb.Empty) (*pb.PrivValidatorKey, error) {
	jsonFile, err := os.Open(a.config.PrivValFile)
	if err != nil {
		return nil, err
	}

	keys := &pb.PrivValidatorKey{}
	err = jsonpb.Unmarshal(jsonFile, keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
