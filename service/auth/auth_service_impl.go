package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/register"
	user "belajar-golang-restful-api/model/web/users"
	repository "belajar-golang-restful-api/repository/auth"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request user.User) (user.Response, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return user.Response{
			Message: err.Error(),
		}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRes, e := service.AuthRepository.CheckLoginValidation(ctx, tx, request)
	helper.PanicIfError(e)

	if userRes.Username != request.Username || userRes.Password != request.Password {
		return user.Response{
			Message: "username atau password salah",
		}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": request.Username,
		"password": request.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(), // token expiry
	})

	tokenString, e := token.SignedString([]byte(middleware.JwtKey))
	helper.PanicIfError(e)

	return user.Response{
		Data:    tokenString,
		Message: "Login Berhasil",
	}, nil

}

func (service *AuthServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserCreateResponse, error) {
	err := service.Validate.Struct(request)
	// helper.PanicIfError(err)
	if err != nil {
		return web.UserCreateResponse{
			Message: err.Error(),
		}, err
	}
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	value, e := service.AuthRepository.CheckDuplicateUser(ctx, tx, request.Username)
	helper.PanicIfError(e)
	if value >= 1 {
		return web.UserCreateResponse{
			Message: "Username already exist",
		}, errors.New("found")
	}

	registerUsers := domain.Register{
		Email:    request.Email,
		Phone:    request.Phone,
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
	}

	registerUsers, err = service.AuthRepository.Save(ctx, tx, registerUsers)

	helper.PanicIfError(err)
	return web.UserCreateResponse{
		Username: registerUsers.Username,
	}, nil
}
