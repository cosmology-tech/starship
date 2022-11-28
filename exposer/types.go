package main

type StatusResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	NodeInfo      NodeInfo      `json:"node_info"`
	ValidatorInfo ValidatorInfo `json:"validator_info"`
}

type NodeInfo struct {
	ID      string `json:"id"`
	Network string `json:"network"`
}

type ValidatorInfo struct {
	Address string `json:"address"`
	PubKey  PubKey `json:"pub_key"`
}

type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
