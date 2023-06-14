package repository

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"oauth2-server/internal/model"
)

type credentialsStore map[string]model.Credentials

type CredentialsStore struct {
	credentialsStore credentialsStore
	logger           *zap.Logger
}

func NewCredentialsStore(l *zap.Logger) CredentialsStore {
	// This can be changed to a database to store all the credentials that the server handles.
	// For this code-challenge, it's just used the following pair username-password as example.

	credentials := credentialsStore{
		"user-test": {
			Username: "user-test",
			Password: "password-test",
		},
	}

	return CredentialsStore{
		credentialsStore: credentials,
		logger:           l,
	}
}

// GetCredentials gets the credentials by username from what it's stored
func (c CredentialsStore) GetCredentials(ctx context.Context, username string) (*model.Credentials, error) {
	if item, found := c.credentialsStore[username]; found {
		return &item, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}
