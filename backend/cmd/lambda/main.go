package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/server"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

func main() {
	ctx := context.Background()
	l := logger.New()

	srv, err := server.New(ctx, l)
	if err != nil {
		l.ErrorContext(ctx, "failed to initialize server", "error", err)
		os.Exit(1)
	}

	go func() {
		if err := srv.Run(ctx); err != nil {
			l.ErrorContext(ctx, "server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful Shutdown (SIGINT / SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Echo の GracefulTimeout (10s) より余裕を持たせ、shutdown 完了を確実に待つ。
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		l.ErrorContext(shutdownCtx, "shutdown error", "error", err)
		os.Exit(1)
	}
}
