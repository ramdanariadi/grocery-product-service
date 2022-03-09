package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-tunas/controllers/category"
	"net/http"
)

func HF(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "tes")
}

func main() {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	impl := category.NewCategoryController(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/category", impl.FindAll)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	errlisten := server.ListenAndServe()
	if errlisten != nil {
		panic("err listen")
	}
}
