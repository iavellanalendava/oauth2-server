package service

import (
	"context"
	"testing"

	"oauth2-server/config"
	"oauth2-server/internal/model"
)

func TestTokenGenerate(t *testing.T) {
	// Create a new instance of ConfigService with mock dependencies
	mockConfig := loadMockConfig()
	service := &ConfigService{
		repoCred: mockCredentialsRepository{},
		config:   mockConfig,
	}

	// Test case: Valid credentials and successful token generation
	clientID := "alice"
	clientSecret := "pass123"
	token, err := service.TokenGenerate(context.Background(), clientID, clientSecret)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if token == nil {
		t.Errorf("expected token not to be nil")
	}

	// Test case: Invalid credentials
	clientID = "eve"
	clientSecret = "pass456"
	token, err = service.TokenGenerate(context.Background(), clientID, clientSecret)
	if err == nil {
		t.Error("expected error, got nil")
	}
	if token != nil {
		t.Errorf("expected token to be nil, got: %v", token)
	}
}

// Mock Credentials Repository implementation for testing
type mockCredentialsRepository struct{}

func (m mockCredentialsRepository) GetCredentials(ctx context.Context, username string) (*model.Credentials, error) {
	// Implement the desired behavior for testing
	if username == "alice" {
		return &model.Credentials{
			Username: "alice",
			Password: "pass123",
		}, nil
	}
	return nil, nil
}

func loadMockConfig() *config.Config {
	cfg := config.Config{}
	cfg.App.Token.Expiration = 3600
	return &cfg
}
