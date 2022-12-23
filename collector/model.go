package main

type ItemsResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
}

func NewItemsResponse(items interface{}) ItemsResponse {
	if items == nil {
		return ItemsResponse{
			Items:      []string{},
			TotalItems: 0,
		}
	}

	listItems, ok := items.([]interface{})
	if !ok {
		return ItemsResponse{
			Items:      []interface{}{items},
			TotalItems: 1,
		}
	}

	return ItemsResponse{
		Items:      listItems,
		TotalItems: len(listItems),
	}
}

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
