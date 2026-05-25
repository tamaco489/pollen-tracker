package main

import (
	"context"
	"os"

	"github.com/tamaco489/pollen-tracker/backend/internal/server"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

func main() {
	ctx := context.Background()
	l := logger.New()
	s := server.New(l)
	l.InfoContext(ctx, "starting server", "addr", ":8080")
	if err := s.Run(ctx); err != nil {
		l.ErrorContext(ctx, "server failed", "error", err)
		os.Exit(1)
	}
}
