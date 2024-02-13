package controller

import "net/http"

type UserController interface {
	CreateUser(wriiter http.ResponseWriter, request *http.Request)
}
