package repository

import (
	"context"
	"fmt"

	"oauth2-server/internal/model"
)

type credentialsStore map[string]model.Credentials

// GetCredentials gets the credentials by username from what it's stored
func (c CredentialsStore) GetCredentials(ctx context.Context, username string) (*model.Credentials, error) {
	if item, found := c.credentialsStore[username]; found {
		return &item, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}
