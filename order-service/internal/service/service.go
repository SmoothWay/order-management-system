package service

import (
	"context"

	"github.com/SmoothWay/oms-order/internal"
)

type OrderService struct {
	storage internal.OrdersStorage
}

func NewService(storage internal.OrdersStorage) *OrderService {
	return &OrderService{
		storage: storage,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context) error {
	return nil
}
