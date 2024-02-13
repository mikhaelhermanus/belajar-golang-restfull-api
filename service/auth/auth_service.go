package service

import (
	web "belajar-golang-restful-api/model/web/register"
	"context"
)

type AuthService interface {
	Create(ctx context.Context, request web.UserCreateRequest) (web.UserCreateResponse, error)
}
