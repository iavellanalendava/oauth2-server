package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"oauth2-server/internal/model"
)

func (c *ConfigController) TokenVerify(ctx *gin.Context) {
	var tokenRequest model.TokenVerificationRequest
	err := ctx.ShouldBindJSON(&tokenRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if tokenRequest.Token == nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to verify token: no token received in request body"))
		return
	}

	tokenData, err := c.service.TokenVerify(ctx, *tokenRequest.Token)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("invalid token: %v", err))
		return
	}

	tokenResponse := model.TokenVerificationResponse{
		AccessToken: *tokenRequest.Token,
		ClientId:    tokenData.Client,
		Scope:       tokenData.Scope,
		ExpiresIn:   tokenData.ExpiresIn,
	}

	ctx.JSON(http.StatusOK, tokenResponse)
}
