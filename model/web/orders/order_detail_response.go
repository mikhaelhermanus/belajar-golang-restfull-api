package web

type OrderStruct struct {
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
}

type OrderDetailResponse struct {
	OrderId  int           `json:"order_id"`
	Products []OrderStruct `json:"products"`
	Total    int           `json:"total`
	Message  string        `json:"message"`
}
