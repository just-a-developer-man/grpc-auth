package main

import (
	"sso/internal/config"
	"sso/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env, cfg.LogLevel)

	// TODO: initialize app
	// TODO: start gRPC-server
}
