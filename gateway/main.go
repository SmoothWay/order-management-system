package main

import (
	"log"
	"net/http"

	"github.com/SmoothWay/oms/commons"
	pb "github.com/SmoothWay/oms/commons/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = commons.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = commons.EnvString("ORDER_SERVICE_ADDR", "localhost:2000")
)

func main() {

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect orderservice: %v", err)
	}
	defer conn.Close()
	log.Println("dialing orders service at", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting http server at %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}
