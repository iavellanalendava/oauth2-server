package repository

import (
	"context"

	"github.com/google/uuid"

	"oauth2-server/internal/model"
)

type CredentialsStoreInterface interface {
	// GetCredentials gets the credentials for a given username
	GetCredentials(ctx context.Context, username string) (*model.Credentials, error)
}

type KeyStoreInterface interface {
	// Save saves a key pair of public and private keys
	Save(key *model.Key) error

	// GetById gets the key pair of public and private keys for a specific user
	GetById(id uuid.UUID) (*model.Key, bool)

	// GetAll gets all the key pairs of public and private keys
	GetAll() []*model.Key
}
