package utils

import (
	"database/sql"
	"fmt"
	"time"
)

func GetDBConnection() *sql.DB {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("open db error")
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
