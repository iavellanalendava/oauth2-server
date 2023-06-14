package service

import (
	"context"

	"oauth2-server/internal/model"
)

func (s *ConfigService) KeysList(ctx context.Context) ([]*model.Key, error) {
	arrKeyList := s.repoKeys.GetAll()

	return arrKeyList, nil
}
