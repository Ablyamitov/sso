package grpcapp

import (
	"fmt"
	"log"
	"net"

	authgrpc "github.com/ablyamiov/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

func New(port int) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)
	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run GRPC Server
func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s:%w", "grpcapp.Run", err)
	}

	log.Println("grpc server is running")
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s:%w", "grpcapp.Run", err)
	}

	return nil
}

// Stop GRPC Server
func (a *App) Stop() {
	log.Println("grpc server is stopping gracefully")
	a.gRPCServer.GracefulStop()
}
