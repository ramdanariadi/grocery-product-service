package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-be-golang/ErrorHandlers"
	"github.com/ramdanariadi/grocery-be-golang/controllers/category"
	"github.com/ramdanariadi/grocery-be-golang/controllers/product"
	"github.com/ramdanariadi/grocery-be-golang/controllers/product/wishlist"
	"github.com/ramdanariadi/grocery-be-golang/controllers/transaction"
	"github.com/ramdanariadi/grocery-be-golang/controllers/transaction/cart"
	"github.com/ramdanariadi/grocery-be-golang/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/helpers"
	"github.com/ramdanariadi/grocery-be-golang/security"
	"github.com/ramdanariadi/grocery-be-golang/utils"
	"log"
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
	wishlistHandler := wishlist.NewWishlistController(db)
	cartHandler := cart.NewCartController(db)
	transactionHandler := transaction.NewTransactionController(db)
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

	router.Handle("/category", security.SecureHandler(categoryHandler.FindAll)).Methods("GET")
	router.Handle("/category/{id}", security.SecureHandler(categoryHandler.FindById)).Methods("GET")
	router.Handle("/category", security.SecureHandler(categoryHandler.Save)).Methods("POST")
	router.Handle("/category/{id}", security.SecureHandler(categoryHandler.Update)).Methods("PUT")
	router.Handle("/category/{id}", security.SecureHandler(categoryHandler.Delete)).Methods("DELETE")

	subrouterProduct := router.PathPrefix("/product").MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		log.Default().Println("mather top product")
		log.Default().Println(match.Vars["id"])
		return true
	}).Subrouter()

	subrouterProduct.Handle("/", security.SecureHandler(productHandler.FindAll)).Methods("GET")
	subrouterProduct.Handle("/category/{id}", security.SecureHandler(productHandler.FindById)).Methods("GET")
	subrouterProduct.Handle("/{id}", security.SecureHandler(productHandler.FindById)).Methods("GET")
	subrouterProduct.Handle("/", security.SecureHandler(productHandler.Save)).Methods("POST")
	subrouterProduct.Handle("/{id}", security.SecureHandler(productHandler.Update)).Methods("PUT")
	subrouterProduct.Handle("/{id}", security.SecureHandler(productHandler.Delete)).Methods("DELETE")

	subrouterTopProduct := router.PathPrefix("/product/top").MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		log.Default().Println("mather top product")
		return true
	}).Subrouter()

	subrouterTopProduct.HandleFunc("/", topProductHandler.FindAll).Methods("GET")
	subrouterTopProduct.HandleFunc("/{id}", topProductHandler.FindById).Methods("GET")
	subrouterTopProduct.HandleFunc("/{id}", topProductHandler.Save).Methods("POST")
	subrouterTopProduct.HandleFunc("/{id}", topProductHandler.Delete).Methods("DELETE")

	subrouterRecommendatinProduct := router.PathPrefix("/product/recommendation").MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool {
		log.Default().Println("mather top product")
		return true
	}).Subrouter()

	subrouterRecommendatinProduct.HandleFunc("/", recommendationProductHandler.FindAll).Methods("GET")
	subrouterRecommendatinProduct.HandleFunc("/{id}", recommendationProductHandler.FindById).Methods("GET")
	subrouterRecommendatinProduct.HandleFunc("/{id}", recommendationProductHandler.Save).Methods("POST")
	subrouterRecommendatinProduct.HandleFunc("/{id}", recommendationProductHandler.Delete).Methods("DELETE")

	subrouterWishlist := router.PathPrefix("/wishlist").Subrouter()
	subrouterWishlist.Handle("/{id}", security.SecureHandler(wishlistHandler.FindByUserId)).Methods("GET")
	subrouterWishlist.Handle("/{userId}/{productId}", security.SecureHandler(wishlistHandler.FindByUserAndProductId)).Methods("GET")
	subrouterWishlist.Handle("/{userId}/{productId}", security.SecureHandler(wishlistHandler.Save)).Methods("POST")
	subrouterWishlist.Handle("/{userId}/{productId}", security.SecureHandler(wishlistHandler.Delete)).Methods("DELETE")

	subrouterCart := router.PathPrefix("/cart").Subrouter()
	subrouterCart.Handle("/{id}", security.SecureHandler(cartHandler.FindById)).Methods("GET")
	subrouterCart.Handle("/{userId}/{productId}/{total:[0-9]+}", security.SecureHandler(cartHandler.Save)).Methods("POST")
	subrouterCart.Handle("/{userId}/{productId}", security.SecureHandler(cartHandler.Delete)).Methods("DELETE")

	subrouterTransaction := router.PathPrefix("/transaction").Subrouter()
	subrouterTransaction.Handle("/", security.SecureHandler(transactionHandler.Save)).Methods("POST")
	subrouterTransaction.Handle("/{id}", security.SecureHandler(transactionHandler.FindById)).Methods("GET")
	subrouterTransaction.Handle("/{id}", security.SecureHandler(transactionHandler.Delete)).Methods("DELETE")
	subrouterTransaction.Handle("/user/{id}", security.SecureHandler(transactionHandler.FindByUserId)).Methods("GET")

	router.Use(ErrorHandlers.PanicHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	errlisten := server.ListenAndServe()
	helpers.PanicIfError(errlisten)
}
