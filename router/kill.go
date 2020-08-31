package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/docker"
)

func kill(ctx *gin.Context) {
	ctx.JSON(200, docker.KillContainers())
	os.Exit(0)
}
