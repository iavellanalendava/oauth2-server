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

	token, err := c.service.TokenVerify(ctx, *tokenRequest.Token)
	if token == nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to verify token: unexpected error"))
		return
	}
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("invalid token: %v", err))
		return
	}

	tokenResponse := model.TokenVerificationResponse{
		Token: tokenRequest.Token,
		Valid: token.Valid,
	}

	ctx.JSON(http.StatusOK, tokenResponse)
}
