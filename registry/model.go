package main

import (
	pb "github.com/Anmol1696/starship/registry/registry"
)

type ChannelPort struct {
	ChannelId string `json:"channel_id"`
	PortId    string `json:"port_id"`
}

type ChannelsInfo struct {
	ChannelPort

	Counterparty ChannelPort
	ConnectionId string
	Ordering     string `json:"ordering,omitempty"`
	Version      string `json:"version,omitempty"`
	State        string `json:"state,omitempty"`
}

type ConnectionClient struct {
	ClientId     string
	ConnectionId string
}

type ConnectionInfo struct {
	ConnectionClient

	Counterparty ConnectionClient
}

type IBCInfo struct {
	ChainId      string `json:"chain_id"`
	ChannelId    string `json:"channel_id"`
	PortId       string `json:"port_id"`
	ConnectionId string `json:"connection_id"`
	ClientId     string `json:"client_id"`
}

type ChainIBCInfo struct {
	IBCInfo

	Counterparty IBCInfo           `json:"counterparty,omitempty"`
	Ordering     string            `json:"ordering,omitempty"`
	Version      string            `json:"version,omitempty"`
	State        string            `json:"state,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
}

func (info *ChainIBCInfo) ToProto() *pb.IBCData {
	return &pb.IBCData{
		Schema: "../ibc_data.schema.json",
		Chain_1: &pb.IBCChain{
			ChainName:    info.ChainId,
			ClientId:     info.ClientId,
			ConnectionId: info.ConnectionId,
		},
		Chain_2: &pb.IBCChain{
			ChainName:    info.Counterparty.ChainId,
			ClientId:     info.Counterparty.ClientId,
			ConnectionId: info.Counterparty.ConnectionId,
		},
		Channels: []*pb.ChannelData{
			&pb.ChannelData{
				Chain_1: &pb.ChannelData_ChannelPort{
					ChannelId: info.ChannelId,
					PortId:    info.PortId,
				},
				Chain_2: &pb.ChannelData_ChannelPort{
					ChannelId: info.Counterparty.ChannelId,
					PortId:    info.Counterparty.PortId,
				},
				Ordering: info.Ordering,
				Version:  info.Version,
				Tags: &pb.ChannelData_Tags{
					// todo: fetch status from client status instead of hardcoding
					Status:    "live",
					Perferred: true,
				},
			},
		},
	}
}

type ChainIBCInfos []*ChainIBCInfo

// GetCounterpartyChainInfo returns ChainIBCInfo struct for given counterparty
// chain id.
func (c ChainIBCInfos) GetCounterpartyChainInfo(chainId string) *ChainIBCInfo {
	for _, cii := range c {
		if cii.Counterparty.ChainId == chainId {
			return cii
		}
	}
	return nil
}
