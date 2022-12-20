package main

type Chain struct {
	Name       string      `json:"name,omitempty"`
	Type       string      `json:"type,omitempty"`
	Validators []Validator `json:"validators,omitempty"`
}

type Validator struct {
	Name    string `json:"name,omitempty"`
	Moniker string `json:"moniker,omitempty"`
	Address string `json:"address,omitempty"`
}

type State struct {
	ID       string `json:"id,omitempty"`
	Height   string `json:"height,omitempty"`
	DataType string `json:"data_type,omitempty"`
}

func NewExportState(id string, height string) State {
	return State{
		ID:       id,
		Height:   height,
		DataType: "json",
	}
}

func NewSnapshotState(id string, height string) State {
	return State{
		ID:       id,
		Height:   height,
		DataType: "tar",
	}
}
