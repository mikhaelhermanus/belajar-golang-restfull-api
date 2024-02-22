package helper

import (
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	productWeb "belajar-golang-restful-api/model/web/produtcs"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToProductResponse(product domain.ProductsAll) productWeb.ProductsResponse {
	return productWeb.ProductsResponse{
		Id:           product.Id,
		Name:         product.Name,
		CategoryName: product.CategoryName,
		Price:        product.Price,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToProductResponses(products []domain.ProductsAll) []productWeb.ProductsResponse {
	var productResponses []productWeb.ProductsResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

// func ToOrderResponses(productOrder []domain.OrderStruct) []orderWeb.OrderDetailResponse{
// 	var productOrderResponses orderWeb.OrderDetailResponse
// 	for _, order := range orders{
// 		orderResponses = append(orderResponses, )
// 	}
// }
