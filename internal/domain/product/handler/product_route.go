package handler

import "fmt"

const PATH = "/products"

func (h *productHandler) MapRoutes() {
	h.mux.HandleFunc(fmt.Sprintf("GET %s", PATH), h.GetProducts())
	h.mux.HandleFunc(fmt.Sprintf("GET %s/{id}", PATH), h.GetProductById())
	h.mux.HandleFunc(fmt.Sprintf("POST %s", PATH), h.CreateProduct())
}
