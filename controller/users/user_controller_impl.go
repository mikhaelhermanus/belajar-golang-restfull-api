package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	webRegister "belajar-golang-restful-api/model/web/register"
	webUser "belajar-golang-restful-api/model/web/users"
	service "belajar-golang-restful-api/service/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type UserControllerImpl struct {
	AuthService service.AuthService
}

func NewUserController(authService service.AuthService) UserController {
	return &UserControllerImpl{
		AuthService: authService,
	}
}

func (controller *UserControllerImpl) LoginUser(writter http.ResponseWriter, request *http.Request) {
	userLoginRequest := webUser.User{}
	err := helper.ReadFromRequestBody(request, &userLoginRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Invalid",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}

	loginUserResponse, e := controller.AuthService.Login(request.Context(), userLoginRequest)

	if e != nil {
		webResponse := web.WebResponse{
			Code:   403,
			Status: "Invalid",
			Data:   loginUserResponse,
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}

	helper.PanicIfError(e)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginUserResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

func (controller *UserControllerImpl) CreateUser(writter http.ResponseWriter, request *http.Request) {
	userCreateRequest := webRegister.UserCreateRequest{}
	err := helper.ReadFromRequestBody(request, &userCreateRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   400,
			Status: "Invalid",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}

	createUserResponse, e := controller.AuthService.Create(request.Context(), userCreateRequest)
	if e != nil {
		webResponse := web.WebResponse{
			Code:   403,
			Status: "Invalid",
			Data:   createUserResponse,
		}
		helper.WriteToResponseBody(writter, webResponse)
		return
	}

	helper.PanicIfError(e)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   createUserResponse,
	}

	helper.WriteToResponseBody(writter, webResponse)
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))

// ValidateMiddleware validates the JWT token.
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return JwtKey, nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(webUser.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(webUser.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(webUser.Exception{Message: "An authorization header is required"})
		}
	})
}
