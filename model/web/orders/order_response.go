package web

type OrderResponse struct {
	OrderId int    `validate:"required" json:"order_id"`
	Message string `json:"message,omitempty"`
}
