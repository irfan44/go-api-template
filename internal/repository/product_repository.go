package repository

import (
	"context"

	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr)
	GetProductById(id int, ctx context.Context) (*entity.Product, errs.MessageErr)
	GenerateProductId(ctx context.Context) int
	CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr)
}

type productRepository struct {
	db []entity.Product
}

func (r *productRepository) GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr) {
	return r.db, nil
}

func (r *productRepository) GetProductById(id int, ctx context.Context) (*entity.Product, errs.MessageErr) {
	for _, product := range r.db {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, errs.NewNotFoundError("Product was not found")
}

func (r *productRepository) GenerateProductId(ctx context.Context) int {
	if len(r.db) == 0 {
		return 1
	}

	return r.db[len(r.db)-1].ID + 1
}

func (r *productRepository) CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr) {
	r.db = append(r.db, product)
	return &product, nil
}

func NewProductRepository(db []entity.Product) ProductRepository {
	return &productRepository{
		db: db,
	}
}
