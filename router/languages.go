package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/config"
)

func languages(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, config.Config.Languages)
}
