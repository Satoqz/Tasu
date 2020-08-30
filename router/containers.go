package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/docker"
)

func containers(ctx *gin.Context) {
	res, err := docker.Run([]string{
		"ps",
		"--filter",
		"name=tasu_",
		"--format",
		"{{.Names}}",
	})
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(res, "\n")
	ctx.JSON(http.StatusOK, list[:len(list)-1])
}
