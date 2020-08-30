package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/config"
	"github.com/satoqz/tasu/docker"
)

type evalRequest struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

func validLanguage(language *string) (result bool) {
	result = false
	for _, val := range config.Config.Languages {
		if val == *language {
			result = true
			return
		}
	}
	return
}

func eval(ctx *gin.Context) {

	var request evalRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing code or language"})
		return
	}

	if !validLanguage(&request.Language) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported language"})
		return
	}

	snowflake := docker.Snowflake.Generate().String()

	log.Printf("Creating unique eval folder in container: tasu_%s with snowflake id: %s\n", request.Language, snowflake)
	res, err := docker.Run([]string{
		"exec",
		fmt.Sprintf("tasu_%s", request.Language),
		"mkdir",
		"-p",
		fmt.Sprintf("eval/%s", snowflake),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Chmod unique eval directory to 777 in container: tasu_%s\n", request.Language)
	res, err = docker.Run([]string{
		"exec",
		fmt.Sprintf("tasu_%s", request.Language),
		"chmod",
		"777",
		fmt.Sprintf("eval/%s", snowflake),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Eval in container: tasu_%s\n", request.Language)
	res, err = docker.Run([]string{
		"exec",
		"-u1001:1001",
		fmt.Sprintf("-w/tmp/eval/%s", snowflake),
		fmt.Sprintf("tasu_%s", request.Language),
		"/bin/sh",
		"/var/run/run.sh",
		request.Code,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"output": res})
}
