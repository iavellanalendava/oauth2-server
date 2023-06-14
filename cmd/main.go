package cmd

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"oauth2-server/config"
	"oauth2-server/internal/controller"
	"oauth2-server/internal/controller/middleware"
	"oauth2-server/repository"
	"oauth2-server/service"
)

func main() {
	// Logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Errorf("failed to load zap logger, error: %v", err))
	}

	// Configuration
	filename := "app-config.yaml"
	cfg, err := config.Load("", filename)
	if err != nil {
		panic(fmt.Errorf("failed to load configuration, error: %v", err))
	}

	// Engine
	engine := gin.Default()
	engine.Use(middleware.RequestLogger(logger))
	engine.Use(middleware.Recovery(logger))
	engine.Use(middleware.ResponseLogger(logger))

	// Repository
	repCred := repository.NewCredentialsStore(logger)
	repKeys := repository.NewKeysStore(logger)

	// Service
	srv := service.NewConfigService(cfg, repCred, repKeys, logger)

	// Controller
	controller.NewConfigController(srv, cfg, engine, logger)

	if err := engine.Run(); err != nil {
		panic(fmt.Errorf("failed to start gin engine, error: %v", err))
	}
}

// database connects to a PostgresSQL database
func database() (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 user=myuser password=mypassword dbname=mydb sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// databaseTable creates the table to store the key pairs (if it doesn't exist)
func databaseTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS key_pairs (
		id SERIAL PRIMARY KEY,
		public_key BYTEA NOT NULL,
		private_key BYTEA NOT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}
