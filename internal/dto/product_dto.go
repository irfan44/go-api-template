package dto

type (
	ProductRequestDTO struct {
		Name        string `json:"name" validate:"required"`
		ProductType string `json:"type" validate:"required"`
	} // @name ProductRequest

	ProductResponseDTO struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ProductType string `json:"type"`
	} // @name ProductResponse

	GetProductsResponseDTO struct {
		BaseResponse
		Data []ProductResponseDTO `json:"data"`
	} // @name GetProductsResponse

	GetProductByIdResponseDTO struct {
		BaseResponse
		Data ProductResponseDTO `json:"data"`
	} // @name GetProductByIdResponse

	CreateProductResponseDTO struct {
		BaseResponse
		Data ProductResponseDTO `json:"data"`
	} // @name CreateProductResponse
)
