package grpcapp

import (
	"fmt"
	authgRPC "grpc-serv/internal/grpc/auth"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// Создаём новый gRPC сервер
func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	// Регистрируем все сервисы
	authgRPC.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// Run — запускает сервер и блокирует выполнение
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// MustRun — вызывает Run и паникует при ошибке
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Stop — корректно останавливает сервер
func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
	a.log.Info("gRPC server stopped")
}
