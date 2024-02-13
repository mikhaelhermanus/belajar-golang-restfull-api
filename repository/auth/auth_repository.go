package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type AuthRepository interface {
	Save(ctx context.Context, tx *sql.Tx, register domain.Register) (domain.Register, error)
}
