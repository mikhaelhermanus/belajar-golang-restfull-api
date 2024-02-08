package service

import (
	web "belajar-golang-restful-api/model/web/produtcs"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductsResponse, error)
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductsResponse
	FindAll(ctx context.Context) []web.ProductsResponse
	FindById(ctx context.Context, productId int) web.ProductsResponse
	Delete(ctx context.Context, productId int)
}
