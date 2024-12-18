package main

import (
	"context"
	"log"
	"net"
  
	"github.com/vitingr/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer() 

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	newGRPCHandler(grpcServer, service)

	service.CreateOrder(context.Background())

	log.Println("GRPC Server Started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}