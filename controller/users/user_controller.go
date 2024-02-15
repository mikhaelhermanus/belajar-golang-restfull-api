package controller

import "net/http"

type UserController interface {
	CreateUser(writter http.ResponseWriter, request *http.Request)
	LoginUser(writter http.ResponseWriter, request *http.Request)
}
