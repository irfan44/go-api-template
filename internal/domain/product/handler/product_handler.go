package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/irfan44/go-api-template/internal/domain/product/service"
	"github.com/irfan44/go-api-template/internal/dto"
	"github.com/irfan44/go-api-template/pkg/errs"
)

type productHandler struct {
	service service.ProductService
	mux     *http.ServeMux
	ctx     context.Context
}

func (h *productHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := h.service.GetProducts(h.ctx)

		if err != nil {
			w.WriteHeader(err.StatusCode())
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(result.ResponseCode)
		json.NewEncoder(w).Encode(result)
	}
}

func (h *productHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.PathValue("id")

		productId, errConv := strconv.Atoi(id)

		if errConv != nil {
			errMsg := errs.NewBadRequest(errConv.Error())
			w.WriteHeader(errMsg.StatusCode())
			json.NewEncoder(w).Encode(errMsg)
			return
		}

		result, errData := h.service.GetProductById(productId, h.ctx)

		if errData != nil {
			w.WriteHeader(errData.StatusCode())
			json.NewEncoder(w).Encode(errData)
			return
		}

		w.WriteHeader(result.ResponseCode)
		json.NewEncoder(w).Encode(result)
	}
}

func (h *productHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		payload := dto.ProductRequestDTO{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			errMsg := errs.NewBadRequest(err.Error())
			w.WriteHeader(errMsg.StatusCode())
			json.NewEncoder(w).Encode(errMsg)
			return
		}

		result, errData := h.service.CreateProduct(payload, h.ctx)

		if errData != nil {
			w.WriteHeader(errData.StatusCode())
			json.NewEncoder(w).Encode(errData)
			return
		}

		w.WriteHeader(result.ResponseCode)
		json.NewEncoder(w).Encode(result)
	}
}

func NewProductHandler(service service.ProductService, mux *http.ServeMux, ctx context.Context) *productHandler {
	return &productHandler{
		service: service,
		mux:     mux,
		ctx:     ctx,
	}
}
