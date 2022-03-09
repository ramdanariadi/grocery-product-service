package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func Test_nil(t *testing.T) {
	var n interface{}
	n = nil
	if n != nil {
		fmt.Println("not nil")
	}
}

func Test_db(t *testing.T) {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from category")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var category string
		var id string
		var image_url interface{}
		var deleted bool

		rowerr := rows.Scan(&category, &id, &deleted, &image_url)

		if rowerr != nil {
			panic(rowerr)
		}

		fmt.Println(image_url)
	}
}
