package app

import (
	"chatroom-demo/internal/app/infrastructure"
	"chatroom-demo/internal/app/infrastructure/inmemory"
	"chatroom-demo/internal/cfg"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Cfg     *cfg.Config
	Storage map[string]infrastructure.IStorage
}

func NewServer(cfg *cfg.Config) *Server {
	server := &Server{
		Cfg:    cfg,
		Engine: gin.Default(),
	}
	setupStorage(server)
	setupRoute(server)
	return server
}

func setupStorage(server *Server) {
	server.Storage = make(map[string]infrastructure.IStorage)
	server.Storage["rooms"] = inmemory.NewStorage()
	server.Storage["users"] = inmemory.NewStorage()
}
