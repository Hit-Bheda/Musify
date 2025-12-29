package http

import (
	"github.com/Hit-Bheda/Musify/internal/http/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthHandler)

	return r
}
