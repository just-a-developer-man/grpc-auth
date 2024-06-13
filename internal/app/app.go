package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCApp *grpcapp.App
}

func NewApp(log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage
	// TODO: init auth service

	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCApp: grpcApp,
	}
}
