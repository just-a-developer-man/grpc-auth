package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/lib/logger"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env, cfg.LogLevel)

	log.Info("starting application", slog.Int("port", cfg.GRPC.Port))

	application := app.NewApp(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCApp.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	recvdSignal := <-stop

	log.Info("stopping application", slog.String("signal", recvdSignal.String()))
	application.GRPCApp.Stop()
	log.Info("application stopped")
}
