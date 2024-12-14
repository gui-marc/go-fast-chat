package main

import (
	"github.com/gui-marc/go-fast-chat/config"
	"github.com/gui-marc/go-fast-chat/internal/app"
)

func main() {
	cfg := config.Load()

	server := app.Init(cfg)

	if err := server.Start(cfg.ServerAddress); err != nil {
		server.Logger.Fatal(err)
	}
}
