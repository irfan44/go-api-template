package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-example/config"
	"github.com/irfan44/go-example/docs"
	"github.com/irfan44/go-example/internal/domain/product/handler"
	"github.com/irfan44/go-example/internal/domain/product/service"
	"github.com/irfan44/go-example/internal/repository"
)

type (
	server struct {
		cfg config.Config
		mux *http.ServeMux
		db  *sql.DB
	}

	repositories struct {
		productRepository repository.ProductRepository
	}

	services struct {
		productService service.ProductService
	}
)

func (s *server) initializeRepositories() *repositories {
	productRepo := repository.NewProductRepository(s.db)

	return &repositories{
		productRepository: productRepo,
	}
}

func (s *server) initializeServices(repo *repositories) *services {
	productService := service.NewProductService(repo.productRepository)

	return &services{
		productService: productService,
	}
}

func (s *server) initializeHandlers(svc *services, v *validator.Validate, ctx context.Context) {
	productHandler := handler.NewProductHandler(svc.productService, s.mux, v, ctx)
	productHandler.MapRoutes()
}

func (s *server) initializeSwagger() {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", s.cfg.Http.Host, s.cfg.Http.Port)
}

func (s *server) initializeServer() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on PORT %s\n", s.cfg.Http.Port)
		if err := s.runHTTPServer(); err != nil {
			log.Printf("Server error: %s\n", err.Error())
		}
	}()

	oscall := <-ch

	fmt.Printf("Server shutdown: %+v\n", oscall)
}

func (s *server) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v := validator.New()

	repo := s.initializeRepositories()
	svc := s.initializeServices(repo)
	s.initializeHandlers(svc, v, ctx)

	s.initializeSwagger()

	s.initializeServer()
}

func NewServer(cfg config.Config, db *sql.DB) *server {
	return &server{
		cfg: cfg,
		mux: http.NewServeMux(),
		db:  db,
	}
}
