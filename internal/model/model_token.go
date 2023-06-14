package model

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
	Status string `json:"status"`
}
