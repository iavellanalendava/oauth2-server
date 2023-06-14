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
	Save(keyPair *model.KeyPair) error

	// GetById gets the key pair of public and private keys for a specific user
	GetById(id int) (*model.KeyPair, error)

	// GetAll gets all the key pairs of public and private keys
	GetAll() ([]*model.KeyPair, error)
}
