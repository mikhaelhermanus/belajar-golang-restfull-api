package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, db *sql.DB, order domain.Orders) (domain.Orders, error)
	CreateOrderDetail(ctx context.Context, tx *sql.Tx, orderDetail domain.OrdersDetail) (domain.OrdersDetail, error)
}
