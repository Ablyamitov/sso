package app

import (
	"time"

	grpcapp "github.com/ablyamiov/sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int, dbURL string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}

}
