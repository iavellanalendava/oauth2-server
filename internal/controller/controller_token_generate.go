package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"oauth2-server/internal/constants"
	"oauth2-server/internal/model"
)

func (c *ConfigController) TokenGenerate(ctx *gin.Context) {
	// Validation
	if ctx.Request.Method != http.MethodPost {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request method"))
		return
	}

	grantType := ctx.Query("grant_type")
	if grantType != "client_credentials" {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid grant type"))
		return
	}

	// Client Credentials
	clientID, clientSecret, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid client credentials"))
		return
	}

	// Token
	signedToken, err := c.service.TokenGenerate(ctx, clientID, clientSecret)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	response := buildTokenResponse(signedToken, strconv.Itoa(c.configuration.App.Token.Expiration))
	ctx.JSON(http.StatusOK, response)
}

func buildTokenResponse(signedToken *string, expiration string) model.TokenResponse {
	return model.TokenResponse{
		AccessToken: *signedToken,
		TokenType:   constants.Bearer,
		Scope:       constants.Read,
		ExpiresIn:   expiration,
	}
}
