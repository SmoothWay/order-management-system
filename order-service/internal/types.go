package internal

import "context"

type OrdersService interface {
	CreateOrder(context.Context) error
}

type OrdersStorage interface {
	Create(context.Context)
}
