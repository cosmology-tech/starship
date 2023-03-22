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

type IBCInfo struct {
	ChainId      string `json:"chain_id"`
	ChannelId    string `json:"channel_id"`
	PortId       string `json:"port_id"`
	ConnectionId string `json:"connection_id"`
	ClientId     string `json:"client_id"`
}

type ChainIBCInfo struct {
	IBCInfo

	CounterParty IBCInfo           `json:"counter_party,omitempty"`
	Ordering     string            `json:"ordering,omitempty"`
	Version      string            `json:"version,omitempty"`
	State        string            `json:"state,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
}

type ChainClient struct {
	logger      *zap.Logger
	chainConfig *lens.ChainClientConfig

	client *lens.ChainClient

	chainIBCInfo []ChainIBCInfo
}

func NewChainClient(logger *zap.Logger, chainID, rpcAddr, home string) (*ChainClient, error) {
	ccc := &lens.ChainClientConfig{
		ChainID:        chainID,
		RPCAddr:        rpcAddr,
		KeyringBackend: "test",
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

// GetIBCInfos will fetch all the IBC channels for the chain
func (c *ChainClient) GetIBCInfos(chainId string) ([]ChainIBCInfo, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	channels, err := querier.Ibc_Channels()
	if err != nil {
		return nil, err
	}

	var chainIBCInfos []ChainIBCInfo
	for _, channel := range channels.Channels {
		fmt.Printf("channel: %s, port: %s\n", channel.ChannelId, channel.PortId)
		fmt.Printf("counter party: channel: %s, port: %s\n", channel.Counterparty.ChannelId, channel.Counterparty.PortId)

		// use connection hop to get connections info
		if len(channel.ConnectionHops) != 1 {
			return nil, fmt.Errorf("number of connections not 1")
		}
		connectionId := channel.ConnectionHops[0]
		clientId, cpConnectionId, cpClientId, err := c.getIBCConnectionInfo(connectionId)
		if err != nil {
			return nil, err
		}

		cpChainId, err := c.getChainIdFromClient(connectionId)
		if err != nil {
			return nil, err
		}

		c := ChainIBCInfo{
			IBCInfo: IBCInfo{
				ChainId:      chainId,
				ChannelId:    channel.ChannelId,
				ConnectionId: connectionId,
				PortId:       channel.PortId,
				ClientId:     clientId,
			},
			CounterParty: IBCInfo{
				ChainId:      cpChainId,
				ChannelId:    channel.Counterparty.ChannelId,
				PortId:       channel.Counterparty.PortId,
				ConnectionId: cpConnectionId,
				ClientId:     cpClientId,
			},
			Ordering: channel.Ordering.String(),
			Version:  channel.Version,
			State:    channel.State.String(),
		}

		chainIBCInfos = append(chainIBCInfos, c)
	}

	return chainIBCInfos, nil
}

func (c *ChainClient) getIBCConnectionInfo(connectionId string) (string, string, string, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	connection, err := querier.Ibc_Connection(connectionId)
	if err != nil {
		return "", "", "", err
	}

	return connection.Connection.ClientId, connection.Connection.Counterparty.ConnectionId, connection.Connection.Counterparty.ClientId, nil
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
