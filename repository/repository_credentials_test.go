package repository

import (
	"context"
	"testing"

	"oauth2-server/internal/model"
)

func TestGetCredentials(t *testing.T) {
	store := CredentialsStore{
		credentialsStore: map[string]model.Credentials{
			"alice": {Username: "alice", Password: "pass123"},
			"bob":   {Username: "bob", Password: "pass456"},
		},
	}

	// Test case: Valid credentials
	username := "alice"
	expectedCredentials := &model.Credentials{Username: "alice", Password: "pass123"}
	credentials, err := store.GetCredentials(context.Background(), username)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if credentials == nil || *credentials != *expectedCredentials {
		t.Errorf("expected credentials: %+v, got: %+v", expectedCredentials, credentials)
	}

	// Test case: Invalid credentials
	username = "eve"
	credentials, err = store.GetCredentials(context.Background(), username)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if credentials != nil {
		t.Errorf("expected credentials to be nil, got: %+v", credentials)
	}
}
