package main

import (
	"context"
	"log"

	pb "github.com/vitingr/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrdersServiceServer

	service OrdersService
}

func newGRPCHandler(grpcServer *grpc.Server, service OrdersService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrdersServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *pb.CreateOrderRequest) (*pb.Order, error) {

	
	log.Println("New order received! Order %v", payload)

	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}