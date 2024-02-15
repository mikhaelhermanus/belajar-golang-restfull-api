package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	web "belajar-golang-restful-api/model/web/users"
	"context"
	"database/sql"
	"errors"
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

func (repository *AuthsRepositoryImpl) CheckLoginValidation(ctx context.Context, tx *sql.Tx, userRequest web.User) (web.User, error) {
	// select username, password from users where username = 'asd';
	// bulk insert
	// insert into product value (?),  value (?),  value (?),  value (?)
	SQL := "select username, password from users where username = ?"

	rows, err := tx.QueryContext(ctx, SQL, userRequest.Username)
	helper.PanicIfError(err)

	defer rows.Close()

	loginUser := web.User{}

	if rows.Next() {
		err := rows.Scan(&loginUser.Username, &loginUser.Password)
		helper.PanicIfError(err)
		return loginUser, nil
	}
	return loginUser, errors.New("Username is not found")
}
