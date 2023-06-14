package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const tokenExpiration = time.Hour

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

	err = savePublicKeyToFile(publicKey, "public_key.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: error saving public key to file: %v\n", err)
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

func savePublicKeyToFile(publicKey *rsa.PublicKey, filename string) error {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, pemBlock)
	if err != nil {
		return err
	}

	return nil
}
