package model

import "github.com/golang-jwt/jwt"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   string `json:"expires_in"`
}

type TokenVerificationRequest struct {
	Token *string `json:"token"`
}

type TokenVerificationResponse struct {
	AccessToken string `json:"access_token"`
	ClientId    string `json:"client_id"`
	Scope       string `json:"scope"`
	ExpiresIn   string `json:"expires_in"`
}

type TokenVerification struct {
	Token     *jwt.Token
	Valid     bool   `json:"valid"`
	Client    string `json:"client_id"`
	Scope     string `json:"scope"`
	ExpiresIn string `json:"expires_in"`
}
