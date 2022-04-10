package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	model2 "github.com/ramdanariadi/grocery-be-golang/main/proto/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)

	listen, err := net.Listen("tcp", "localhost:9000")
	helpers.PanicIfError(err)

	grpcServer := grpc.NewServer()

	implementedServer := model2.NewProductServiceServerImpl(db)
	model2.RegisterProductServiceServer(grpcServer, implementedServer)

	err = grpcServer.Serve(listen)
	helpers.PanicIfError(err)
	log.Println("gRPC server running on port 9000")
}
