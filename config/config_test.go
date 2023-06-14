package config

import (
	"testing"
)

func TestLoadConfigTokenExpiration(t *testing.T) {
	testPath := "../test/"
	testFilename := "test-config.yaml"

	config, err := Load(testPath, testFilename)
	if err != nil {
		t.Fatal(err)
	}

	expectedToken := Token{
		Expiration: 10,
	}

	if config.App.Token.Expiration != expectedToken.Expiration {
		t.Errorf("unexpected token configuration: expected: %v, got: %v", expectedToken, config.App.Token.Expiration)
	}
}
