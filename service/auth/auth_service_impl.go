package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/register"
	repository "belajar-golang-restful-api/repository/auth"
	"context"
	"database/sql"

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

func (service *AuthServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) (web.UserCreateResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

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
