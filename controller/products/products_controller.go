package controller

import (
	"net/http"
)

type ProductController interface {
	Create(writter http.ResponseWriter, request *http.Request)
	FindAll(writter http.ResponseWriter, request *http.Request)
	FindById(writter http.ResponseWriter, request *http.Request)
	Update(writter http.ResponseWriter, request *http.Request)
	Delete(writter http.ResponseWriter, request *http.Request)
}
