package app

import (
	"fmt"
	"log/slog"
	"time"

	grpcapp "grpc-serv/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
	log     *slog.Logger
	port    int
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		log:     log,
		GRPCSrv: grpcApp,
		port:    grpcPort,
	}
}

// MustRun запускает все сервисы приложения
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run запускает gRPC сервер
func (a *App) Run() error {
	const op = "app.Run"
	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	log.Info("starting services")

	// Запуск gRPC сервера через метод grpcapp.App
	if err := a.GRPCSrv.Run(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop останавливает все сервисы приложения
func (a *App) Stop() {
	const op = "app.Stop"
	a.log.With(slog.String("op", op)).Info("stopping services", slog.Int("port", a.port))

	a.GRPCSrv.Stop()
}
