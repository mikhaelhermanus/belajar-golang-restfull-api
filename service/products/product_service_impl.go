package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/produtcs"
	repository "belajar-golang-restful-api/repository/products"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductsRepository repository.ProductsRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProductService(productRepository repository.ProductsRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductsRepository: productRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductsResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := domain.Products{
		Name:       request.Name,
		CategoryId: request.CategoryId,
	}

	products, err = service.ProductsRepository.Save(ctx, tx, products)
	helper.PanicIfError(err)
	return web.ProductsResponse{
		Id:   products.Id,
		Name: products.Name,
	}, nil
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	products := service.ProductsRepository.FindAll(ctx, tx)
	return helper.ToProductResponses(products)
}
