package repository

import (
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/users"
	"context"
	"database/sql"
)

type AuthRepository interface {
	Save(ctx context.Context, tx *sql.Tx, register domain.Register) (domain.Register, error)
	CheckDuplicateUser(ctx context.Context, tx *sql.Tx, userName string) (value int, e error)
	CheckLoginValidation(ctx context.Context, tx *sql.Tx, userRequest web.User) (web.User, error)
}
