package main

import (
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	protoModel "github.com/ramdanariadi/grocery-be-golang/main/proto/model"
	"github.com/ramdanariadi/grocery-be-golang/main/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := utils.NewDbConnection()

	listen, err := net.Listen("tcp", "localhost:9000")
	helpers.PanicIfError(err)

	grpcServer := grpc.NewServer()

	implementedServer := protoModel.NewProductServiceServerImpl(db)
	protoModel.RegisterProductServiceServer(grpcServer, implementedServer)

	err = grpcServer.Serve(listen)
	helpers.PanicIfError(err)
	log.Println("gRPC server running on port 9000")
}
