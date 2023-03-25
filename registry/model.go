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

func NewChainRegistry(info ChainIBCInfo) pb.IBCData {
	return pb.IBCData{}
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
