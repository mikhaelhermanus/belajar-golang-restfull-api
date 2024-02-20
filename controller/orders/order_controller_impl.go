package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	webOrder "belajar-golang-restful-api/model/web/orders"
	service "belajar-golang-restful-api/service/orders"
	"log"
	"net/http"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(writter http.ResponseWriter, request *http.Request) {
	orderCreateRequest := webOrder.OrderCreateRequest{}
	err := helper.ReadFromRequestBody(request, &orderCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Invalid",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}

	orderResponse, e := controller.OrderService.CreateOrder(request.Context(), orderCreateRequest)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	if e != nil {
		log.Println(e, "line order controller 43")
	}
	helper.WriteToResponseBody(writter, webReponse)
}
