package domain

type OrderStruct struct {
	ProductId int
	Price     int
	Quantity  int
}

type OrdersDetail struct {
	Products []OrderStruct
	OrderId  int
}
