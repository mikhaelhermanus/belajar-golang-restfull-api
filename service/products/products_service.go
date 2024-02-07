package service

import (
	web "belajar-golang-restful-api/model/web/produtcs"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductsResponse, error)
	FindAll(ctx context.Context) []web.ProductsResponse
	FindById(ctx context.Context, productId int) web.ProductsResponse
}
