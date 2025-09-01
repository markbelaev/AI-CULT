package server

import (
	"AI-CULT/internal/server/routes"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() *Server {
	router := gin.Default()

	server := &Server{
		router: router,
	}

	routes.SetupRoutes(router)

	return server
}

func (s *Server) Start() error {
	slog.Info("Starting server")
	return s.router.Run(":8080")
}
