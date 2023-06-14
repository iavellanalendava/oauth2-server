package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"oauth2-server/config"
	"oauth2-server/internal/controller/paths"
	"oauth2-server/service"
)

type ConfigController struct {
	service       service.ConfigServiceInterface
	configuration *config.Config
	engine        *gin.Engine
	logger        *zap.Logger
}

func NewConfigController(s service.ConfigServiceInterface, c *config.Config, e *gin.Engine, l *zap.Logger) {
	controllerConfig := &ConfigController{
		service:       s,
		configuration: c,
		engine:        e,
		logger:        l,
	}
	controllerConfig.setUpRoutes()
}

func (c *ConfigController) setUpRoutes() {
	c.engine.POST(paths.TokenGenerate, c.TokenGenerate)
	c.engine.POST(paths.TokenVerify, c.TokenVerify)
	c.engine.GET(paths.KeysList, c.KeysList)
}
