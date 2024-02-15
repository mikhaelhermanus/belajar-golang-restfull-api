package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	webProducts "belajar-golang-restful-api/model/web/produtcs"
	service "belajar-golang-restful-api/service/products"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writter http.ResponseWriter, request *http.Request) {
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
		log.Println(e, "line 45")
	}
	helper.WriteToResponseBody(writter, webResponse)

}

func (controller *ProductControllerImpl) FindById(writter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productId := vars["productId"]

	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request) {
	productResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *ProductControllerImpl) Update(writter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productUpdateRequest := webProducts.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := vars["productId"]
	id, err := strconv.Atoi(productId) // conversi string ke object 'Atoi'
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *ProductControllerImpl) Delete(writter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	productId := vars["productId"]
	id, err := strconv.Atoi(productId) // conversi string ke object 'Atoi'
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writter, webResponse)
}
