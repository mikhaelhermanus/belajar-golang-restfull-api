package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"log"
)

type AuthsRepositoryImpl struct {
}

func NewAuthsRepository() AuthRepository {
	return &AuthsRepositoryImpl{}
}

func (repository *AuthsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, register domain.Register) (domain.Register, error) {
	SQL := "insert into users (email, phone, username, password, name) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, register.Email, register.Phone, register.Username, register.Password, register.Name)
	if err != nil {
		log.Println(err.Error(), "line 21")
		return domain.Register{}, err
	}
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	register.Id = int(id)
	return register, nil
}

func (repository *AuthsRepositoryImpl) CheckDuplicateUser(ctx context.Context, tx *sql.Tx, userName string) (value int, e error) {
	SQL := "select count(username) from users where username = ?"

	e = tx.QueryRowContext(ctx, SQL, userName).Scan(&value)

	return value, e
}
