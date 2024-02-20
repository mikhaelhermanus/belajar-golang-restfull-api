package web

type ProductStruct struct {
	ProductId int `validate:"required" json:"product_id"`
	Price     int `validate:"required" json:"price"`
	Quantity  int `validate:"required" json:"quantity"`
}

type OrderCreateRequest struct {
	Product []ProductStruct `json:"products"`
}
