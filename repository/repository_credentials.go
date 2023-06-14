package repository

import (
	"context"

	"oauth2-server/internal/model"
)

type memoryStore map[string]model.Credentials

// GetCredentials gets the credentials by username from what it's stored
func (c CredentialsStore) GetCredentials(ctx context.Context, username string) (*model.Credentials, error) {
	if item, found := c.memoryStore[username]; found {
		return &item, nil
	}

	return nil, nil
}
