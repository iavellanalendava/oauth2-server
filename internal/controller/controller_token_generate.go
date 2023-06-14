package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"oauth2-server/internal/constants"
	"oauth2-server/internal/model"
)

func (c *ConfigController) TokenGenerate(ctx *gin.Context) {
	clientId, _ := ctx.Get(constants.ClientId)
	clientSecret, _ := ctx.Get(constants.ClientSecret)

	signedToken, err := c.service.TokenGenerate(ctx, clientId.(string), clientSecret.(string))
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
