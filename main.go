package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"go-tunas/ErrorHandlers"
	"go-tunas/controllers/category"
	"go-tunas/helpers"
	"go-tunas/security"
	"net/http"
)

func main() {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)

	categoryHandler := category.NewCategoryController(db)
	securityHandler := security.NewSecurityController(db)

	router := httprouter.New()

	router.POST("/login", securityHandler.Login)
	router.POST("/register", securityHandler.SignUp)

	router.GET("/category", security.SecureHandler(categoryHandler.FindAll))
	router.GET("/category/:id", security.SecureHandler(categoryHandler.FindById))
	router.POST("/category", security.SecureHandler(categoryHandler.Save))
	router.PUT("/category/:id", security.SecureHandler(categoryHandler.Update))
	router.DELETE("/category/:id", security.SecureHandler(categoryHandler.Delete))

	router.PanicHandler = ErrorHandlers.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	errlisten := server.ListenAndServe()
	helpers.PanicIfError(errlisten)
}
