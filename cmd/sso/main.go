package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ablyamiov/sso/internal/app"
	"github.com/ablyamiov/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log.Println("Starting application")

	application := app.New(cfg.GRPC.Port, cfg.DB.URL, cfg.TokenTTL)

	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GRPCServer.Stop()
	log.Println("Application stopped")
}
