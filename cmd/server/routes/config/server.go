package config

import (
	"github.com/gin-gonic/gin"
	"log"
)

type ConfigurationServer struct {
	port   string
	server *gin.Engine
}

func NewServer() ConfigurationServer {
	return ConfigurationServer{
		port:   "8000",
		server: gin.Default(),
	}
}

func (s *ConfigurationServer) Run() {
	router := ConfigurationRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
