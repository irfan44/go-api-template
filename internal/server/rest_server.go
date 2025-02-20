package server

import (
	"fmt"
	"net/http"

	"github.com/irfan44/go-api-template/internal/domain/product/handler"
	"github.com/irfan44/go-api-template/internal/domain/product/service"
	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/internal/repository"
)

const PORT = ":8080"

func Run() {
	productRepo := repository.NewProductRepository([]entity.Product{})

	productService := service.NewProductService(productRepo)

	mux := http.NewServeMux()

	productHandler := handler.NewProductHandler(productService, mux)

	productHandler.MapRoutes()

	fmt.Printf("Listening on PORT %s\n", PORT)
	http.ListenAndServe(PORT, mux)
}
