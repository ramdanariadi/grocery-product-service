package main

import (
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/service/cart"
	"github.com/ramdanariadi/grocery-product-service/main/service/category"
	"github.com/ramdanariadi/grocery-product-service/main/service/product"
	"github.com/ramdanariadi/grocery-product-service/main/service/transaction"
	"github.com/ramdanariadi/grocery-product-service/main/service/wishlist"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := utils.NewDbConnection()

	listen, err := net.Listen("tcp", "localhost:9000")
	helpers.PanicIfError(err)

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

	log.Println("gRPC server running on port 9000")

	err = grpcServer.Serve(listen)
	helpers.PanicIfError(err)
}
