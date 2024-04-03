package service

import (
	"context"
	"product/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context, page, limit int) ([]domain.Product, int, error) {
	products, totalData, err := s.repo.GetAll(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return products, totalData, nil
}

func (s *ProductService) Create(ctx context.Context, product *domain.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
	return s.repo.GetById(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, id primitive.ObjectID, product *domain.Product) error {
	return s.repo.Update(ctx, id, product)
}

func (s *ProductService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
