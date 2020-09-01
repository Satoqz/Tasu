package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/containers"
)

func status(ctx *gin.Context) {
	containers.RefreshMap()
	ctx.JSON(http.StatusOK, containers.Map)
}
