package main

type PrivValKey struct {
	Addr   string `json:"address,omitempty"`
	PubKey Key    `json:"pub_key,omitempty"`
	Priv   Key    `json:"priv_key,omitempty"`
}

type Key struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
