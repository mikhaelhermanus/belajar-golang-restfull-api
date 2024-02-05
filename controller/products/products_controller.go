package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
