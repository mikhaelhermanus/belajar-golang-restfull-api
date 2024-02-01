package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	//something (context, realtion transicition, params what we need) return something
	//ctx => tipe data context.
	// 1. find duplicate (name string) (value int, e error)
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
