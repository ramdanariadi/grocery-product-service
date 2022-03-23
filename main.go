package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go-tunas/controllers/category"
	"go-tunas/controllers/product"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	"go-tunas/security"
	"go-tunas/utils"
	"net/http"
)

func main() {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)

	categoryHandler := category.NewCategoryController(db)
	productHandler := product.NewProductControllerImpl(db)
	topProductHandler := product.NewTopProductControllerImpl(db)
	recommendationProductHandler := product.NewRcmdProductControllerImpl(db)
	securityHandler := security.NewSecurityController(db)

	router := mux.NewRouter()

	router.HandleFunc("/login", securityHandler.Login).Methods("POST")
	router.HandleFunc("/register", securityHandler.SignUp).Methods("POST")

	router.HandleFunc("/readcsv/:filename", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		filename := vars["filename"]
		customresponses.SendResponse(writer, utils.ProductsFromCSV("others/"+filename), 200)
	}).Methods("GET")

	router.HandleFunc("/ws", utils.WS).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/category", security.SecureHandler(categoryHandler.FindAll)).Methods("GET")
	router.HandleFunc("/category/:id", security.SecureHandler(categoryHandler.FindById)).Methods("GET")
	router.HandleFunc("/category", security.SecureHandler(categoryHandler.Save)).Methods("POST")
	router.HandleFunc("/category/:id", security.SecureHandler(categoryHandler.Update)).Methods("PUT")
	router.HandleFunc("/category/:id", security.SecureHandler(categoryHandler.Delete)).Methods("DELETE")

	router.HandleFunc("/product", security.SecureHandler(productHandler.FindAll)).Methods("GET")
	router.HandleFunc("/product/category/:id", security.SecureHandler(productHandler.FindById)).Methods("GET")
	router.HandleFunc("/product/:id", security.SecureHandler(productHandler.FindById)).Methods("GET")
	router.HandleFunc("/product", security.SecureHandler(productHandler.Save)).Methods("POST")
	router.HandleFunc("/product/:id", security.SecureHandler(productHandler.Update)).Methods("PUT")
	router.HandleFunc("/product/:id", security.SecureHandler(productHandler.Delete)).Methods("DELETE")

	router.HandleFunc("/product/top", topProductHandler.FindAll).Methods("GET")
	router.HandleFunc("/product/top/:id", topProductHandler.FindById).Methods("GET")
	router.HandleFunc("/product/top", topProductHandler.Save).Methods("POST")
	router.HandleFunc("/product/top/:id", topProductHandler.Update).Methods("PUT")
	router.HandleFunc("/product/top/:id", topProductHandler.Delete).Methods("DELETE")

	router.HandleFunc("/product/recommendation", recommendationProductHandler.FindAll).Methods("GET")
	router.HandleFunc("/product/recommendation/:id", recommendationProductHandler.FindById).Methods("GET")
	router.HandleFunc("/product/recommendation", recommendationProductHandler.Save).Methods("POST")
	router.HandleFunc("/product/recommendation/:id", recommendationProductHandler.Update).Methods("PUT")
	router.HandleFunc("/product/recommendation/:id", recommendationProductHandler.Delete).Methods("DELETE")

	//router.PanicHandler = ErrorHandlers.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	errlisten := server.ListenAndServe()
	helpers.PanicIfError(errlisten)
}
