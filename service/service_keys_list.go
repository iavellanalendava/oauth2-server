package service

import (
	"context"

	"oauth2-server/internal/model"
)

func (s *ConfigService) KeysList(ctx context.Context, clientId string) ([]*model.Key, error) {
	arrKeyList, err := s.repoKeys.GetAllById(clientId)
	if err != nil {
		return nil, err
	}

	return arrKeyList, nil
}
