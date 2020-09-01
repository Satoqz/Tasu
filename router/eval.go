package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/satoqz/tasu/config"
	"github.com/satoqz/tasu/containers"
)

type evalRequest struct {
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

func validLanguage(language *string) (result bool) {
	result = false
	for _, val := range config.Languages {
		if val == *language {
			result = true
			return
		}
	}
	return
}

func createNode() (node *snowflake.Node) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	return
}

var node = createNode()

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

	containerName := fmt.Sprintf("tasu_%s", request.Language)

	container := containers.Map[containerName]

	if container.Language == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Container not found"})
		return
	}

	if !container.Alive {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Currently waiting for container restart"})
		return
	}

	snowflake := node.Generate().String()

	log.Printf("Creating unique eval folder in container: tasu_%s with snowflake id: %s\n", request.Language, snowflake)
	res, err := containers.Run([]string{
		"exec",
		containerName,
		"mkdir",
		"-p",
		fmt.Sprintf("eval/%s", snowflake),
	})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Container currently unavailable"})
		return
	}

	log.Printf("Chmod unique eval directory to 777 in container: %s\n", containerName)
	res, err = containers.Run([]string{
		"exec",
		containerName,
		"chmod",
		"777",
		fmt.Sprintf("eval/%s", snowflake),
	})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Container currently unavailable"})
		return
	}

	log.Printf("Eval in container: %s\n", containerName)
	res, err = containers.Run([]string{
		"exec",
		"-u1001:1001",
		fmt.Sprintf("-w/tmp/eval/%s", snowflake),
		containerName,
		"/bin/sh",
		"/var/run/run.sh",
		request.Code,
	})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "Container currently unavailable"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"output": res})
}
