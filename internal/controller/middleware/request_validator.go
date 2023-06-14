package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		ctx.Set("clientID", clientID)
		ctx.Set("clientSecret", clientSecret)
		ctx.Next()
	}
}
