package router

import (
	"github.com/gin-gonic/gin"
)

// Setup creates a new gin server and attaches all routes to it
func Setup() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// GET routes
	r.GET("/", root)
	r.GET("languages", languages)
	r.GET("/containers", containers)
	r.GET("/shutdown", shutdown)
	r.GET("/cleanup", cleanup)
	// POST routes
	r.POST("/eval", eval)

	r.Run()
}
