package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/irfan44/go-api-template/config"
	"github.com/irfan44/go-api-template/internal/domain/product/handler"
	"github.com/irfan44/go-api-template/internal/domain/product/service"
	"github.com/irfan44/go-api-template/internal/entity"
	"github.com/irfan44/go-api-template/internal/repository"
)

const PORT = ":8080"

type server struct {
	cfg config.Config
	mux *http.ServeMux
}

func (s *server) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	productRepo := repository.NewProductRepository([]entity.Product{})

	productService := service.NewProductService(productRepo)

	productHandler := handler.NewProductHandler(productService, s.mux, ctx)

	productHandler.MapRoutes()

	go func() {
		log.Printf("Listening on PORT: %s\n", s.cfg.Http.Port)
		if err := s.runHTTPServer(); err != nil {
			log.Printf("s.runHTTPServer: %s\n", err.Error())
		}
	}()

	oscall := <-ch

	fmt.Printf("system call: %+v\n", oscall)
}

func NewServer(cfg config.Config) *server {
	return &server{
		cfg: cfg,
		mux: http.NewServeMux(),
	}
}
