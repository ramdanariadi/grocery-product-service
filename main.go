package main

import (
	_ "github.com/lib/pq"
	cart2 "github.com/ramdanariadi/grocery-product-service/main/cart"
	category2 "github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	product2 "github.com/ramdanariadi/grocery-product-service/main/product"
	transaction2 "github.com/ramdanariadi/grocery-product-service/main/transaction"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	wishlist2 "github.com/ramdanariadi/grocery-product-service/main/wishlist"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := utils.NewDbConnection()

	listen, err := net.Listen("tcp", ":50051")
	helpers.PanicIfError(err)

	grpcServer := grpc.NewServer()

	productImplementedServer := product2.NewProductServiceServerImpl(db)
	product2.RegisterProductServiceServer(grpcServer, productImplementedServer)

	categoryImplementedServer := category2.NewCategoryServiceServerImpl(db)
	category2.RegisterCategoryServiceServer(grpcServer, categoryImplementedServer)

	cartImplementedServer := cart2.NewCartServiceImpl(db)
	cart2.RegisterCartServiceServer(grpcServer, cartImplementedServer)

	wishlistImplementedServer := wishlist2.NewWishlistServer(db)
	wishlist2.RegisterWishlistServiceServer(grpcServer, wishlistImplementedServer)

	transactionImplementedServer := transaction2.NewTransactionServiceServer(db)
	transaction2.RegisterTransactionServiceServer(grpcServer, transactionImplementedServer)

	log.Println("gRPC server running on port 50051")

	err = grpcServer.Serve(listen)
	helpers.PanicIfError(err)
}
