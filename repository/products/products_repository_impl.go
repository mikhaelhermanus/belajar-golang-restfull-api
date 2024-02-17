package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
	"log"
)

type ProductsRepositoryImpl struct {
}

func NewProductsRepository() ProductsRepository {
	return &ProductsRepositoryImpl{}
}

func (repository *ProductsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, products domain.Products) (domain.Products, error) {
	SQL := "insert into products (name, category_id, price) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, products.Name, products.CategoryId, products.Price)
	if err != nil {
		log.Println(err.Error(), "line 22")
		return domain.Products{}, err
	}
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	products.Id = int(id)
	return products, nil
}

func (repository *ProductsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.ProductsAll, error) {
	// select products.id as product_id, products.name as product_name, category.name as category from products join category on category.id = products.category_id where products.id = 2;
	SQL := "select products.id as product_id, products.name as product_name, category.name as category from products join category on category.id = products.category_id where products.id = ?"

	rows, err := tx.QueryContext(ctx, SQL, productId)

	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.ProductsAll{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.CategoryName)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("Product Not Found")
	}
}

func (repository *ProductsRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductsAll {
	// select products.id as product_id, products.name as product_name, category.name as category from products join category on category.id = products.category_id;
	SQL := "select products.id as product_id, products.name as product_name, category.name as category from products join category on category.id = products.category_id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.ProductsAll
	for rows.Next() {
		product := domain.ProductsAll{}
		// scan ordering by sql syntax we fill example sql : id, name, categoryName
		err := rows.Scan(&product.Id, &product.Name, &product.CategoryName)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (repository *ProductsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Products) error {
	SQL := "update products set name = ? , category_id = ? where id = ?"
	// _ tidak perlu handle return value use _
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return nil
}

func (respository *ProductsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	SQL := "delete from products where id = ?"

	_, err := tx.ExecContext(ctx, SQL, productId)

	helper.PanicIfError(err)
}
