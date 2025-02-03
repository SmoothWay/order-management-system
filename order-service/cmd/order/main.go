package main

import (
	"context"

	"github.com/SmoothWay/oms-order/internal/service"
	"github.com/SmoothWay/oms-order/internal/storage"
)

func main() {

	storage := storage.NewStorage()
	svc := service.NewService(storage)

	svc.CreateOrder(context.Background())
}
