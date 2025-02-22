package handler

import (
	"net/http"

	"github.com/irfan44/go-example/pkg/internal_http"
)

func (h *productHandler) MapRoutes() {
	h.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodGet, "/products"),
		h.GetProducts(),
	)
	h.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodGet, "/products/{id}"),
		h.GetProductById(),
	)
	h.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodPost, "/products"),
		h.CreateProduct(),
	)
	h.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodPut, "/products/{id}"),
		h.UpdateProduct(),
	)
}
