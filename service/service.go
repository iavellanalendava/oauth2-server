package service

import (
	"go.uber.org/zap"

	"oauth2-server/config"
	"oauth2-server/repository"
)

type ConfigService struct {
	config   *config.Config
	repoCred repository.CredentialsStore
	repoKeys repository.KeysStore
	logger   *zap.Logger
}

func NewConfigService(cfg *config.Config, rC repository.CredentialsStore, rK repository.KeysStore, l *zap.Logger) ConfigServiceInterface {
	return &ConfigService{
		config:   cfg,
		repoCred: rC,
		repoKeys: rK,
		logger:   l,
	}
}
