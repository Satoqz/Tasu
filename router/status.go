package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/containers"
)

func status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, containers.Map)
}