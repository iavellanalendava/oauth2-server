package service

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (s *ConfigService) TokenVerify(ctx context.Context, token string) (*jwt.Token, error) {
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

	return t, nil
}
