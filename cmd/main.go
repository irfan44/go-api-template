package main

import (
	"log"

	"github.com/irfan44/go-example/config"
	"github.com/irfan44/go-example/internal/dependency"
	"github.com/irfan44/go-example/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("err loading .env file: %s\n", err.Error())
	}
}

// @title Example Go REST API
// @version 1.0
// @description Basic REST API using Golang
// @BasePath /
func main() {
	cfg := config.NewConfig()

	db, err := dependency.InitializeDB(cfg)
	if err != nil {
		return
	}

	s := server.NewServer(cfg, db)

	s.Run()
}
