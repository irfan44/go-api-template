package repository

import (
	"context"
	"database/sql"

	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr)
	GetProductById(id int, ctx context.Context) (*entity.Product, errs.MessageErr)
	CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr)
}

type productRepository struct {
	db *sql.DB
}

func (r *productRepository) GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr) {
	query := `
		SELECT id, name, producttype FROM product
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, errs.NewBadRequest("Cannot get products")
	}

	result := []entity.Product{}

	for rows.Next() {
		product := entity.Product{}

		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.ProductType,
		); err != nil {
			return nil, errs.NewUnprocessibleEntityError("Failed to parse data")
		}

		result = append(result, product)
	}

	return result, nil
}

func (r *productRepository) GetProductById(id int, ctx context.Context) (*entity.Product, errs.MessageErr) {
	query := `
		SELECT id, name, producttype FROM product WHERE id = $1
	`

	product := entity.Product{}

	if err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.ProductType,
	); err != nil {
		return nil, errs.NewNotFoundError("Product was not found")
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr) {
	query := `
		INSERT INTO product (name, producttype) VALUES ($1, $2) RETURNING id, name, producttype;
	`

	newProduct := entity.Product{}

	if err := r.db.QueryRow(query, product.Name, product.ProductType).Scan(
		&newProduct.ID,
		&newProduct.Name,
		&newProduct.ProductType,
	); err != nil {
		return nil, errs.NewBadRequest("Cannot add product")
	}

	return &newProduct, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
