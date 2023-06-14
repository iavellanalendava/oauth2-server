package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"oauth2-server/internal/model"
)

func (s *ConfigService) TokenGenerate(ctx context.Context, clientId, clientSecret string) (*string, error) {
	// Authentication
	credentials, err := s.repoCred.GetCredentials(ctx, clientId)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	if credentials == nil {
		return nil, fmt.Errorf("failed to generate token: no user was found by clientId")
	}

	if credentials.Password != clientSecret {
		return nil, fmt.Errorf("failed to generate token: clientSecret does not match with the one from the request")
	}

	// Private and Public Keys
	privateKey, publicKey, errKeys := generateRSAKeyPair(2048)
	if errKeys != nil {
		return nil, fmt.Errorf("failed to generate token: error generating RSA key pair: %v\n", errKeys)
	}

	id := uuid.New()
	err = s.savePublicKey(credentials.Username, id, publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: error saving public key: %v", err)
	}

	// Token
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = credentials.Username
	claims["iss"] = "oauth2-server"
	claims["aud"] = "audience-test"
	claims["exp"] = s.config.App.Token.Expiration
	claims["its"] = time.Now().Unix()

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: error in signing the token")
	}

	return &signedToken, nil
}

func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate RSA key pair: %v", err)
	}
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func (s *ConfigService) savePublicKey(clientId string, id uuid.UUID, publicKey *rsa.PublicKey) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	key := &model.Key{
		Id:        id,
		ClientId:  clientId,
		PublicKey: publicKeyBytes,
		CreatedAt: time.Now(),
	}

	if err = s.repoKeys.Save(key); err != nil {
		return err
	}

	return nil
}
