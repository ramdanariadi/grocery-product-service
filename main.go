package main

import (
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/cart"
	cartModel "github.com/ramdanariadi/grocery-product-service/main/cart/model"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	categoryModel "github.com/ramdanariadi/grocery-product-service/main/category/model"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/transaction"
	transactionModel "github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist"
	wishlistModel "github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	utils.PanicIfError(err)
	err = db.AutoMigrate(&categoryModel.Category{}, &model.Product{}, &wishlistModel.Wishlist{}, &cartModel.Cart{}, &transactionModel.Transaction{}, &transactionModel.TransactionDetail{})
	utils.LogIfError(err)
	listen, err := net.Listen("tcp", ":50051")
	utils.PanicIfError(err)

	grpcServer := grpc.NewServer()

	productImplementedServer := product.NewProductServiceServerImpl(db)
	product.RegisterProductServiceServer(grpcServer, productImplementedServer)

	categoryImplementedServer := category.NewCategoryServiceServerImpl(db)
	category.RegisterCategoryServiceServer(grpcServer, categoryImplementedServer)

	cartImplementedServer := cart.NewCartServiceImpl(db)
	cart.RegisterCartServiceServer(grpcServer, cartImplementedServer)

	wishlistImplementedServer := wishlist.NewWishlistServer(db)
	wishlist.RegisterWishlistServiceServer(grpcServer, wishlistImplementedServer)

	transactionImplementedServer := transaction.NewTransactionServiceServer(db)
	transaction.RegisterTransactionServiceServer(grpcServer, transactionImplementedServer)

	log.Println("gRPC server running on port 50051")

	err = grpcServer.Serve(listen)
	utils.PanicIfError(err)
}
