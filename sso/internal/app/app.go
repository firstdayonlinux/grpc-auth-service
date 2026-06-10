package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	authgrpc "sso/internal/grpc/auth"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
	authService authgrpc.Auth,

) *App {
	grpcApp := grpcapp.New(log, grpcPort, authService)

	return &App{
		GRPCSrv: grpcApp,
	}
}
