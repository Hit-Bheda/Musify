// Package http contains app api related data like server and routes etc
package http

import (
	"github.com/Hit-Bheda/Musify/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthHandler)
	r.POST("/add-song", handlers.AddSong)

	return r
}
