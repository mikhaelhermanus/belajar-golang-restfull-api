package controller

import "net/http"

type OrderController interface {
	Create(writter http.ResponseWriter, request *http.Request)
	FindById(writter http.ResponseWriter, request *http.Request)
}
