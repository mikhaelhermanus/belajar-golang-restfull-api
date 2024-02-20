package service

import (
	web "belajar-golang-restful-api/model/web/orders"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request web.OrderCreateRequest) (web.OrderResponse, error)
}
