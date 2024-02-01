package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"log"
	"time"
)

func NewDB(local string, database string) *sql.DB {
	connection := "root:" + local + "/" + database
	db, err := sql.Open("mysql", connection)
	helper.PanicIfError(err)
	log.Println("running on user :", local)
	log.Println("running on database:", database)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
