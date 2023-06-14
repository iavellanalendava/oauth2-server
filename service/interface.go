package service

import (
	"context"

	"oauth2-server/internal/model"
)

type ConfigServiceInterface interface {
	// TokenGenerate generates a new access token
	TokenGenerate(ctx context.Context, clientId, clientSecret string) (*string, error)

	// TokenVerify verifies the status of a token
	TokenVerify(ctx context.Context, token string) (model.TokenVerification, error)

	// KeysList retrieves the list of signing keys stored
	KeysList(ctx context.Context) ([]*model.Key, error)
}
