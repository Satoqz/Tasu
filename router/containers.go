package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/docker"
)

func containers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, docker.Containers)
}
