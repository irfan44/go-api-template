package repository

import (
	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type ProductRepository interface {
	GetProducts() ([]entity.Product, errs.MessageErr)
	GetProductById(id int) (*entity.Product, errs.MessageErr)
	CreateProduct(product entity.Product) (*entity.Product, errs.MessageErr)
}

type productRepository struct {
	db []entity.Product
}

func (r *productRepository) GetProducts() ([]entity.Product, errs.MessageErr) {
	return r.db, nil
}

func (r *productRepository) GetProductById(id int) (*entity.Product, errs.MessageErr) {
	for _, product := range r.db {
		if product.ID == id {
			return &product, nil
		}
	}

	return nil, errs.NewNotFoundError("Product was not found")
}

func (r *productRepository) CreateProduct(product entity.Product) (*entity.Product, errs.MessageErr) {
	r.db = append(r.db, product)
	return &product, nil
}

func NewProductRepository(db []entity.Product) ProductRepository {
	return &productRepository{
		db: db,
	}
}
