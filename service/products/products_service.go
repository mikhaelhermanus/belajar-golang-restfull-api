package service

import (
	web "belajar-golang-restful-api/model/web/produtcs"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductsResponse, error)
}
