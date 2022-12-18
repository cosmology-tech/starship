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
	PubKey  Key    `json:"pub_key"`
}

type PrivValKey struct {
	Address string `json:"address,omitempty"`
	PubKey  Key    `json:"pub_key,omitempty"`
	Priv    Key    `json:"priv_key,omitempty"`
}

type Key struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
