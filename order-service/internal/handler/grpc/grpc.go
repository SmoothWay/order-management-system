package grpc

import (
	"context"
	"log"

	"github.com/SmoothWay/oms-order/internal/service"
	pb "github.com/SmoothWay/oms/commons/api"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service *service.OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, service *service.OrderService) {
	pb.RegisterOrderServiceServer(
		grpcServer, &grpcHandler{service: service},
	)

}

func (g *grpcHandler) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received")
	o := &pb.Order{
		ID:         uuid.NewString(),
		CustomerID: request.CustomerID,
		Status:     "OK",
	}
	return o, nil
}
