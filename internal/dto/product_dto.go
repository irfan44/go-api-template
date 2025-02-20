package dto

type ProductRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	ProductType string `json:"type" validate:"required"`
}

type ProductResponseDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProductType string `json:"type"`
}

type GetProductsResponseDTO struct {
	BaseResponse
	Data []ProductResponseDTO `json:"data"`
}

type GetProductByIdResponseDTO struct {
	BaseResponse
	Data ProductResponseDTO `json:"data"`
}

type CreateProductResponseDTO struct {
	BaseResponse
	Data ProductResponseDTO `json:"data"`
}
