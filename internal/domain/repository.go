package domain

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	GetAll(ctx context.Context, page int, limit int) ([]Product, int, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, id primitive.ObjectID, product *Product) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
