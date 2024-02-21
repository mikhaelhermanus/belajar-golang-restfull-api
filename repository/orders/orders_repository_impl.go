package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"log"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (repository *OrdersRepositoryImpl) CreateOrder(ctx context.Context, db *sql.DB, order domain.Orders) (domain.Orders, error) {
	SQL := "insert into orders(total) values (?)"
	result, err := db.ExecContext(ctx, SQL, order.Total)
	if err != nil {
		log.Println(err.Error(), "line 21 order rep")
		return domain.Orders{}, err
	}
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.OrderId = int(id)
	return order, nil
}

func (repository *OrdersRepositoryImpl) CreateOrderDetail(ctx context.Context, tx *sql.Tx, orderDetail domain.OrdersDetail) (domain.OrdersDetail, error) {
	log.Println(orderDetail.Products, "line 33")
	SQL := "insert into orders_detail(id_product, id_order, price, quantity) Values (?, ?, ?, ?)"

	for _, row := range orderDetail.Products {
		_, err := tx.ExecContext(ctx, SQL, row.ProductId, orderDetail.OrderId, row.Price, row.Quantity)
		if err != nil {
			log.Println(err.Error(), "line 22")
			return domain.OrdersDetail{}, err
		}

	}

	return orderDetail, nil

}
