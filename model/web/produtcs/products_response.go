package web

type ProductsResponse struct {
	// if isempty message , will return undefined
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
	Price        int    `json:"price,omitempty"`
	Message      string `json:"message,omitempty"`
}
