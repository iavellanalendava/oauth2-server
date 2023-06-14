package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt"

	"oauth2-server/internal/model"
)

func (s *ConfigService) TokenVerify(ctx context.Context, token string) (*model.TokenVerification, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return nil, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	client, expiration, scope, err := extractTokenInfo(t)
	tokenData := &model.TokenVerification{
		Token:     t,
		Valid:     t.Valid,
		Client:    client,
		Scope:     scope,
		ExpiresIn: strconv.FormatInt(expiration, 10),
	}

	return tokenData, nil
}

func extractTokenInfo(token *jwt.Token) (string, int64, string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, "", fmt.Errorf("failed to extract token claims")
	}

	clientID, ok := claims["client"].(string)
	if !ok {
		return "", 0, "", fmt.Errorf("failed to extract client ID from token claims")
	}

	expiration, ok := claims["exp"].(int64)
	if !ok {
		return "", 0, "", fmt.Errorf("failed to extract expiration time from token claims")
	}

	scope, ok := claims["scope"].(string)
	if !ok {
		return "", 0, "", fmt.Errorf("failed to extract scope from token claims")
	}

	return clientID, expiration, scope, nil
}
