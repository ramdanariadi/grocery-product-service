package main

import (
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/cart"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/transaction"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))

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
