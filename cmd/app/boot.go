package app

import (
	"fmt"
	"chatroom-demo/internal/app"
	"chatroom-demo/internal/cfg"
)

func Bootstrap() {
	server := app.NewServer(&cfg.Cfg)
	server.Run(fmt.Sprintf(":%s", server.Cfg.App.ServerPort))
}
