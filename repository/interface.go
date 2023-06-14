package repository

import (
	"context"

	"oauth2-server/internal/model"
)

type CredentialsStoreInterface interface {
	// GetCredentials gets the credentials for a given username
	GetCredentials(ctx context.Context, username string) (*model.Credentials, error)
}

type KeyStoreInterface interface {
	// Save saves a key pair of public and private keys
	Save(clientId string, key *model.Key) error

	// GetAllById gets all the keys for a specific clientId
	GetAllById(clientId string) ([]*model.Key, error)
}
