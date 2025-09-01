package server

import (
	"log/slog"
	"net/http"

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

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	s.router.GET("/health", s.healthHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.Request.Context()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (s *Server) Start() error {
	slog.Info("Starting server")
	return s.router.Run(":8080")
}
