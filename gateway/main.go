package main

import (
	"log"
	"net/http"

	pb "github.com/vitingr/common/api"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vitingr/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":3000")
	orderServiceAddr = "localhost:2000"
)

func main() {
	connection, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial server: %v", err)
	}
	defer connection.Close()

	log.Println("Dialing orders service at ", orderServiceAddr)

	c := pb.NewOrdersServiceClient(connection)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
