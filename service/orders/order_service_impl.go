package service

import (
	"belajar-golang-restful-api/exception"
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
	total := 0
	arrayProduct := []domain.OrderStruct{}

	for _, item := range request.Product {
		total += item.Price * item.Quantity
		arrayProduct = append(arrayProduct, domain.OrderStruct{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	orders := domain.Orders{
		Total: total,
	}

	orders, err = service.OrderRepository.CreateOrder(ctx, service.DB, orders)
	helper.PanicIfError(err)

	ordersDetail := domain.OrdersDetail{
		Products: arrayProduct,
		OrderId:  orders.OrderId,
	}

	ordersDetail, err = service.OrderRepository.CreateOrderDetail(ctx, tx, ordersDetail)
	log.Println(err, "line 56")
	helper.PanicIfError(err)
	return web.OrderResponse{
		OrderId: ordersDetail.OrderId,
		Message: "Order Berhasil terbentuk",
	}, nil

}

func (service *OrderServiceImpl) FindById(ctx context.Context, orderId int) web.OrderDetailResponse {
	order, err := service.OrderRepository.FindById(ctx, service.DB, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var arrayProduct []web.OrderStruct
	var total = 0
	for _, item := range order.Products {
		arrayProduct = append(arrayProduct, web.OrderStruct{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Name:      item.Name,
		})
		total += item.Quantity * item.Price
	}

	return web.OrderDetailResponse{
		Products: arrayProduct,
		Total:    total,
		OrderId:  order.OrderId,
		Message:  "Success",
	}
}
