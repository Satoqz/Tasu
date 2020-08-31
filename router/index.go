package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.String(
		http.StatusOK,
		"Welcome to Tasū!",
	)
}
