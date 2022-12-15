package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"time"
)

func NewDbConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:secret@localhost/grocery-product-service?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, err
}
