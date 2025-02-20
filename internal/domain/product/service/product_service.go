package service

import (
	"net/http"

	"github.com/irfan44/go-api-template/internal/dto"
	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/internal/repository"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type ProductService interface {
	GetProducts() (*dto.GetProductsResponseDTO, errs.MessageErr)
	GetProductById(id int) (*dto.GetProductByIdResponseDTO, errs.MessageErr)
	CreateProduct(product dto.ProductRequestDTO) (*dto.CreateProductResponseDTO, errs.MessageErr)
}

type productService struct {
	repository repository.ProductRepository
}

func (s *productService) GetProducts() (*dto.GetProductsResponseDTO, errs.MessageErr) {
	products, err := s.repository.GetProducts()

	if err != nil {
		return nil, err
	}

	result := dto.GetProductsResponseDTO{
		BaseResponse: dto.BaseResponse{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "SUCCESS",
		},
		Data: entity.Products(products).ToProductsDTO(),
	}

	return &result, nil
}

func (s *productService) GetProductById(id int) (*dto.GetProductByIdResponseDTO, errs.MessageErr) {
	product, err := s.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	result := dto.GetProductByIdResponseDTO{
		BaseResponse: dto.BaseResponse{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "SUCCESS",
		},
		Data: *product.ToProductResponseDTO(),
	}

	return &result, nil
}

func (s *productService) CreateProduct(product dto.ProductRequestDTO) (*dto.CreateProductResponseDTO, errs.MessageErr) {
	newProductId := s.repository.GenerateProductId()

	newProductEntity := entity.Product{
		ID:          newProductId,
		Name:        product.Name,
		ProductType: product.ProductType,
	}

	newProduct, err := s.repository.CreateProduct(newProductEntity)

	if err != nil {
		return nil, err
	}

	result := dto.CreateProductResponseDTO{
		BaseResponse: dto.BaseResponse{
			ResponseCode:    http.StatusOK,
			ResponseMessage: "SUCCESS",
		},
		Data: *newProduct.ToProductResponseDTO(),
	}

	return &result, nil
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}
