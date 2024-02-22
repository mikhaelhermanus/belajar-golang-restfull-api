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

func (repository *OrdersRepositoryImpl) FindById(ctx context.Context, db *sql.DB, orderId int) (domain.OrdersDetail, error) {
	sql := "SELECT products.id as id_products, products.name, orders_detail.quantity, orders_detail.price from orders JOIN orders_detail ON (orders_detail.id_order = orders.id) JOIN products ON (products.id = orders_detail.id_product ) where orders.id = ?"

	rows, err := db.QueryContext(ctx, sql, orderId)

	helper.PanicIfError(err)

	defer rows.Close()

	log.Println(rows, "line 60")

	var products []domain.OrderStruct

	for rows.Next() {
		product := domain.OrderStruct{}
		err := rows.Scan(&product.ProductId, &product.Name, &product.Quantity, &product.Price)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return domain.OrdersDetail{
		Products: products,
		OrderId:  orderId,
	}, nil
}
