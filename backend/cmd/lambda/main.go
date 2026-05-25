package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/handler"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

func main() {
	ctx := context.Background()

	l := logger.New()

	e := echo.New()

	h := handler.New(l)
	gen.RegisterHandlers(e, gen.NewStrictHandler(h, nil))

	l.InfoContext(ctx, "starting server", "addr", ":8080")
	if err := e.Start(":8080"); err != nil {
		l.ErrorContext(ctx, "server failed", "error", err)
		os.Exit(1)
	}
}
