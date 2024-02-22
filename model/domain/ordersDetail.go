package domain

type OrderStruct struct {
	ProductId int
	Price     int
	Quantity  int
	Name      string
}

type OrdersDetail struct {
	Products []OrderStruct
	OrderId  int
}
