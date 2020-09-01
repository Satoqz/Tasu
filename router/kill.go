package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/containers"
)

func kill(ctx *gin.Context) {
	ctx.JSON(200, containers.Kill())
}
