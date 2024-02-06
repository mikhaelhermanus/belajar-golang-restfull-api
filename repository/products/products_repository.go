package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type ProductsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, products domain.Products) (domain.Products, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductsAll
}
