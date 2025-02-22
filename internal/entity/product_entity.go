package entity

import "github.com/irfan44/go-example/internal/dto"

type Product struct {
	ID          int
	Name        string
	ProductType string
}

func (e *Product) ToProductResponseDTO() *dto.ProductResponseDTO {
	return &dto.ProductResponseDTO{
		ID:          e.ID,
		Name:        e.Name,
		ProductType: e.ProductType,
	}
}

type Products []Product

func (e Products) ToProductsDTO() []dto.ProductResponseDTO {
	result := []dto.ProductResponseDTO{}

	for _, product := range e {
		result = append(result, *product.ToProductResponseDTO())
	}

	return result
}
