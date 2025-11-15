package main

import (
	"log/slog"
	"os"

	"github.com/pixambi/simple-server/internal/config"
	"github.com/pixambi/simple-server/internal/server"
)

func main() {
	cfg := config.Load()

	slog.SetDefault(cfg.Logger)

	server := server.New(cfg)
	if err := server.Start(true); err != nil {
		slog.Error("Failed to run server", "error", err)
		os.Exit(1)
	}
}
