package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
)

func NewDbConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:secret@localhost/grocery-product-service?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)
	return db, err
}
