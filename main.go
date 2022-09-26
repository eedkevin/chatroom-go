package main

import (
	"chatroom-demo/cmd/app"
	"chatroom-demo/internal/cfg"
)

func main() {
	cfg.Init()
	app.Bootstrap()
}
