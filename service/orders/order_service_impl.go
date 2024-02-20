package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/orders"
	repository "belajar-golang-restful-api/repository/orders"
	"context"
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrdersRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrdersRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request web.OrderCreateRequest) (web.OrderResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.OrderResponse{
			Message: err.Error(),
		}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := domain.Orders{
		Total: 50000,
	}

	orders, err = service.OrderRepository.CreateOrder(ctx, tx, orders)
	helper.PanicIfError(err)

	ordersDetail := domain.OrdersDetail{
		Products: []domain.OrderStruct{
			domain.OrderStruct{ProductId: 15, Quantity: 2, Price: 250000},
		},
		OrderId: orders.OrderId,
	}

	ordersDetail, err = service.OrderRepository.CreateOrderDetail(ctx, tx, ordersDetail)
	log.Println(err, "line 56")
	helper.PanicIfError(err)
	return web.OrderResponse{
		OrderId: ordersDetail.OrderId,
		Message: "Order Berhasil terbentuk",
	}, nil

}
