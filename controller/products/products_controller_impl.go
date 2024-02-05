package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	webProducts "belajar-golang-restful-api/model/web/produtcs"
	service "belajar-golang-restful-api/service/products"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := webProducts.ProductCreateRequest{}
	err := helper.ReadFromRequestBody(request, &productCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Invalid",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}
	productResponse, e := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	if e != nil {
		log.Println(e)
	}
	helper.WriteToResponseBody(writter, webResponse)

}
