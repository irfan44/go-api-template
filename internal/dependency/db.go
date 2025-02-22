package dependency

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/irfan44/go-example/config"
	"github.com/irfan44/go-example/pkg/postgres"
)

func InitializeDB(cfg config.Config) (*sql.DB, error) {
	db, err := postgres.NewDB(
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
		return nil, err
	}

	fmt.Println("Connect to database successful")

	return db, nil
}
