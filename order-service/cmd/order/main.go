package main

import (
	"context"
	"log"
	"net"

	"github.com/SmoothWay/commons"
	grpchandler "github.com/SmoothWay/oms-order/internal/handler/grpc"
	"github.com/SmoothWay/oms-order/internal/service"
	"github.com/SmoothWay/oms-order/internal/storage"
	"google.golang.org/grpc"
)

var (
	grpcAddr = commons.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed")
	}
	defer l.Close()
	storage := storage.NewStorage()
	svc := service.NewService(storage)

	svc.CreateOrder(context.Background())

	grpchandler.NewGRPCHandler(grpcServer, svc)
	log.Printf("starting grpc server on port %v\n", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
