package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type ProductsRepositoryImpl struct {
}

func NewProductsRepository() ProductsRepository {
	return &ProductsRepositoryImpl{}
}

func (repository *ProductsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, products domain.Products) (domain.Products, error) {
	SQL := "insert into products (name, category_id) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, products.Name, products.CategoryId)
	if err != nil {
		return domain.Products{}, err
	}
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	products.Id = int(id)
	return products, nil
}
