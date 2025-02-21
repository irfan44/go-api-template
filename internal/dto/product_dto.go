package dto

type ProductRequestDTO struct {
	Name        string `json:"name" validate:"required"`
	ProductType string `json:"type" validate:"required"`
} // @name ProductRequest

type ProductResponseDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProductType string `json:"type"`
}

type GetProductsResponseDTO struct {
	BaseResponse
	Data []ProductResponseDTO `json:"data"`
} // @name GetProductsResponse

type GetProductByIdResponseDTO struct {
	BaseResponse
	Data ProductResponseDTO `json:"data"`
} // @name GetProductByIdResponse

type CreateProductResponseDTO struct {
	BaseResponse
	Data ProductResponseDTO `json:"data"`
} // @name CreateProductResponse
