package main

type ItemsResponse struct {
	Items      interface{} `json:"items"`
	TotalItems int         `json:"total_items"`
}

func NewItemsResponse(items interface{}) ItemsResponse {
	if items == nil {
		return ItemsResponse{[]string{}, 0}
	}

	listItems, ok := items.([]interface{})
	if !ok {
		return ItemsResponse{[]interface{}{items}, 1}
	}

	return ItemsResponse{listItems, len(listItems)}
}
