package main

import (
	"encoding/json"
	"fmt"
	"os"

	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	lens "github.com/strangelove-ventures/lens/client"
	"github.com/strangelove-ventures/lens/client/query"
	"go.uber.org/zap"
)

type ChainClient struct {
	logger      *zap.Logger
	chainConfig *lens.ChainClientConfig

	client *lens.ChainClient
}

func NewChainClient(logger *zap.Logger, chainID, rpcAddr, home string) (*ChainClient, error) {
	ccc := &lens.ChainClientConfig{
		ChainID:        chainID,
		RPCAddr:        rpcAddr,
		KeyringBackend: "test",
		Key:            "default",
		//RPCAddr:        "https://osmosis-1.technofractal.com:443",
		//GRPCAddr:       "https://gprc.osmosis-1.technofractal.com:443",
		//AccountPrefix:  "osmo",
		//GasAdjustment:  1.2,
		//GasPrices:      "0.01uosmo",
		//MinGasAmount:   0,
		Debug:   true,
		Timeout: "20s",
		//OutputFormat: "json",
		SignModeStr: "direct",
	}
	client, err := lens.NewChainClient(logger, ccc, home, os.Stdin, os.Stdout)
	if err != nil {
		return nil, err
	}

	chainClient := &ChainClient{
		logger:      logger,
		chainConfig: ccc,
		client:      client,
	}
	return chainClient, nil
}

// getChannelPort returns the chains and the counterparty info
func (c *ChainClient) getChannelsPorts() ([]ChannelsInfo, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	channels, err := querier.Ibc_Channels()
	if err != nil {
		return nil, err
	}

	var channelsInfo []ChannelsInfo
	for _, channel := range channels.Channels {
		// use connection hop to get connections info
		if len(channel.ConnectionHops) != 1 {
			return nil, fmt.Errorf("number of connections not 1")
		}

		ci := ChannelsInfo{
			ChannelPort: ChannelPort{
				ChannelId: channel.ChannelId,
				PortId:    channel.PortId,
			},
			Counterparty: ChannelPort{
				ChannelId: channel.Counterparty.ChannelId,
				PortId:    channel.Counterparty.PortId,
			},
			ConnectionId: channel.ConnectionHops[0],
			Ordering:     channel.Ordering.String(),
			Version:      channel.Version,
		}

		channelsInfo = append(channelsInfo, ci)
	}

	return channelsInfo, nil
}

func (c *ChainClient) getConnectionClient(connectionId string) (*ConnectionInfo, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	connection, err := querier.Ibc_Connection(connectionId)
	if err != nil {
		return nil, err
	}

	return &ConnectionInfo{
		ConnectionClient: ConnectionClient{
			ConnectionId: connectionId,
			ClientId:     connection.Connection.ClientId,
		},
		Counterparty: ConnectionClient{
			ConnectionId: connection.Connection.Counterparty.ConnectionId,
			ClientId:     connection.Connection.Counterparty.ClientId,
		},
	}, nil
}

// GetIBCClients will fetch all the IBC channels for the chain
func (c *ChainClient) getChainIdFromClient(clientId string) (string, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	state, err := querier.Ibc_ClientState(clientId)
	if err != nil {
		return "", err
	}

	clientState, err := clienttypes.UnpackClientState(state.ClientState)
	if err != nil {
		return "", err
	}

	if clientState.ClientType() != exported.Tendermint {
		return "", fmt.Errorf("client state type not %s", exported.Tendermint)
	}

	cs := &ibctm.ClientState{}
	err = json.Unmarshal(state.ClientState.Value, cs)
	if err != nil {
		return "", err
	}

	return cs.ChainId, nil
}

// GetIBCInfos will fetch all the IBC channels for the chain
func (c *ChainClient) GetIBCInfos() ([]ChainIBCInfo, error) {
	channelsInfo, err := c.getChannelsPorts()
	if err != nil {
		return nil, err
	}

	var chainIBCInfos []ChainIBCInfo
	for _, channelInfo := range channelsInfo {
		cpChainId, err := c.getChainIdFromClient(channelInfo.ConnectionId)
		if err != nil {
			return nil, err
		}

		connectionInfo, err := c.getConnectionClient(channelInfo.ConnectionId)
		if err != nil {
			return nil, err
		}

		cii := ChainIBCInfo{
			IBCInfo: IBCInfo{
				ChainId:      c.chainConfig.ChainID,
				ChannelId:    channelInfo.ChannelId,
				PortId:       channelInfo.PortId,
				ConnectionId: connectionInfo.ConnectionId,
				ClientId:     connectionInfo.ClientId,
			},
			Counterparty: IBCInfo{
				ChainId:      cpChainId,
				ChannelId:    channelInfo.Counterparty.ChannelId,
				PortId:       channelInfo.Counterparty.PortId,
				ConnectionId: connectionInfo.Counterparty.ConnectionId,
				ClientId:     connectionInfo.Counterparty.ClientId,
			},
			Ordering: channelInfo.Ordering,
			Version:  channelInfo.Version,
			State:    channelInfo.State,
		}

		chainIBCInfos = append(chainIBCInfos, cii)
	}

	return chainIBCInfos, nil
}
