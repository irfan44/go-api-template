package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type (
	ProductRepository interface {
		GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr)
		GetProductById(id int, ctx context.Context) (*entity.Product, errs.MessageErr)
		CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr)
		UpdateProduct(product entity.Product, id int, ctx context.Context) (*entity.Product, errs.MessageErr)
	}

	productRepository struct {
		db *sql.DB
	}
)

func (r *productRepository) GetProducts(ctx context.Context) ([]entity.Product, errs.MessageErr) {
	query := `
		SELECT id, name, producttype FROM product
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, errs.NewInternalServerError()
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

	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.ProductType,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError(
				fmt.Sprintf("Product with id %d not found", id),
			)
		}
		return nil, errs.NewInternalServerError()
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(product entity.Product, ctx context.Context) (*entity.Product, errs.MessageErr) {
	query := `
		INSERT INTO product (name, producttype) VALUES ($1, $2) RETURNING id, name, producttype;
	`

	newProduct := entity.Product{}

	if err := r.db.QueryRowContext(ctx, query, product.Name, product.ProductType).Scan(
		&newProduct.ID,
		&newProduct.Name,
		&newProduct.ProductType,
	); err != nil {
		return nil, errs.NewInternalServerError()
	}

	return &newProduct, nil
}

func (r *productRepository) UpdateProduct(product entity.Product, id int, ctx context.Context) (*entity.Product, errs.MessageErr) {
	query := `
		UPDATE product SET name = $1, producttype = $2 WHERE id = $3 RETURNING id, name, producttype;
	`

	updatedProduct := entity.Product{}

	if err := r.db.QueryRowContext(ctx, query, product.Name, product.ProductType, id).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.ProductType,
	); err != nil {
		return nil, errs.NewInternalServerError()
	}

	return &updatedProduct, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
