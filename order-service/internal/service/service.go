package service

import (
	"context"

	"github.com/SmoothWay/oms-order/internal"
)

type Service struct {
	storage internal.OrdersStorage
}

func NewService(storage internal.OrdersStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateOrder(ctx context.Context) error {
	return nil
}
