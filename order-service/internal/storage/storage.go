package storage

import "context"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (S *Storage) Create(ctx context.Context) {

}
