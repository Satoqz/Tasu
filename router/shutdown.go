package router

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/docker"
)

func shutdown(ctx *gin.Context) {
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
	list = list[:len(list)-1]
	for _, item := range list {
		_, err := docker.Run([]string{
			"kill",
			item,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx.JSON(200, list)
}
