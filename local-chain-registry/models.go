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
