package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"go-tunas/ErrorHandlers"
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

	router := httprouter.New()

	router.POST("/login", securityHandler.Login)
	router.POST("/register", securityHandler.SignUp)

	router.GET("/readcsv/:filename", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		filename := params.ByName("filename")
		customresponses.SendResponse(writer, utils.ProductsFromCSV("others/"+filename), 200)
	})

	router.GET("/ws", utils.WS)

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET("/category", security.SecureHandler(categoryHandler.FindAll))
	router.GET("/category/:id", security.SecureHandler(categoryHandler.FindById))
	router.POST("/category", security.SecureHandler(categoryHandler.Save))
	router.PUT("/category/:id", security.SecureHandler(categoryHandler.Update))
	router.DELETE("/category/:id", security.SecureHandler(categoryHandler.Delete))

	router.GET("/product", security.SecureHandler(productHandler.FindAll))
	router.GET("/product/:id", security.SecureHandler(productHandler.FindById))
	router.GET("/product/category/:id", security.SecureHandler(productHandler.FindById))
	router.POST("/product", security.SecureHandler(productHandler.Save))
	router.PUT("/product/:id", security.SecureHandler(productHandler.Update))
	router.DELETE("/product/:id", security.SecureHandler(productHandler.Delete))

	router.GET("/product/top", topProductHandler.FindAll)
	router.GET("/product/top/:id", topProductHandler.FindById)
	router.POST("/product/top", topProductHandler.Save)
	router.PUT("/product/top/:id", topProductHandler.Update)
	router.DELETE("/product/top/:id", topProductHandler.Delete)

	router.GET("/product/recommendation", recommendationProductHandler.FindAll)
	router.GET("/product/recommendation/:id", recommendationProductHandler.FindById)
	router.POST("/product/recommendation", recommendationProductHandler.Save)
	router.PUT("/product/recommendation/:id", recommendationProductHandler.Update)
	router.DELETE("/product/recommendation/:id", recommendationProductHandler.Delete)

	router.PanicHandler = ErrorHandlers.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	errlisten := server.ListenAndServe()
	helpers.PanicIfError(errlisten)
}
