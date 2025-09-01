package routes

import (
	"AI-CULT/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	healthHandler := handlers.NewHealthHandler()

	router.GET("/health", healthHandler.Health)
}
