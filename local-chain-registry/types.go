package main

type ItemsResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
}

func NewItemsResponse(items ...interface{}) ItemsResponse {
	return ItemsResponse{
		Items:      items,
		TotalItems: len(items),
	}
}
