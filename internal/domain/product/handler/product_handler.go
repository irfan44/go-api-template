package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-example/internal/domain/product/service"
	"github.com/irfan44/go-example/internal/dto"
	"github.com/irfan44/go-example/pkg/errs"
	"github.com/irfan44/go-example/pkg/internal_http"
)

type productHandler struct {
	service service.ProductService
	mux     *http.ServeMux
	v       *validator.Validate
	ctx     context.Context
}

// @Summary Get All Products
// @Tags products
// @Produce json
// @Success 200 {object} GetProductsResponse
// @Router /products [get]
func (h *productHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.service.GetProducts(h.ctx)

		if err != nil {
			internal_http.SendResponse(w, err.StatusCode(), err)
			return
		}

		internal_http.SendResponse(w, result.ResponseCode, result)
	}
}

// @Summary Get Products by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} GetProductByIdResponse
// @Router /products/{id} [get]
func (h *productHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		productId, errConv := strconv.Atoi(id)

		if errConv != nil {
			errMsg := errs.NewBadRequest(errConv.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		result, errData := h.service.GetProductById(productId, h.ctx)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.ResponseCode, result)
	}
}

// @Summary Create New Product
// @Tags products
// @Accept json
// @Produce json
// @Param requestBody body ProductRequest true "Request Body"
// @Success 200 {object} CreateProductResponse
// @Router /products [post]
func (h *productHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := dto.ProductRequestDTO{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			errMsg := errs.NewBadRequest(err.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		errVal := h.v.Struct(payload)

		if errVal != nil {
			errMsg := errs.NewBadRequest(errVal.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		result, errData := h.service.CreateProduct(payload, h.ctx)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.ResponseCode, result)
	}
}

// @Summary Update Product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param requestBody body ProductRequest true "Request Body"
// @Success 200 {object} UpdateProductResponse
// @Router /products/{id} [put]
func (h *productHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		productId, errConv := strconv.Atoi(id)

		if errConv != nil {
			errMsg := errs.NewBadRequest(errConv.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		payload := dto.ProductRequestDTO{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			errMsg := errs.NewBadRequest(err.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		errVal := h.v.Struct(payload)

		if errVal != nil {
			errMsg := errs.NewBadRequest(errVal.Error())
			internal_http.SendResponse(w, errMsg.StatusCode(), errMsg)
			return
		}

		result, errData := h.service.UpdateProduct(payload, productId, h.ctx)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.ResponseCode, result)
	}
}

func NewProductHandler(service service.ProductService, mux *http.ServeMux, v *validator.Validate, ctx context.Context) *productHandler {
	return &productHandler{
		service: service,
		mux:     mux,
		v:       v,
		ctx:     ctx,
	}
}
