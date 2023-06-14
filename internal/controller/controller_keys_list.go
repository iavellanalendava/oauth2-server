package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *ConfigController) KeysList(ctx *gin.Context) {
	keysList, err := c.service.KeysList(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, keysList)
}
