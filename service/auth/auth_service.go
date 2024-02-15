package service

import (
	web "belajar-golang-restful-api/model/web/register"
	user "belajar-golang-restful-api/model/web/users"
	"context"
)

type AuthService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserCreateResponse, error)
	Login(ctx context.Context, request user.User) (user.Response, error)
}
