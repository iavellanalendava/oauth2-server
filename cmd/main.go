package cmd

import (
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
