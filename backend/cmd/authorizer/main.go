package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/tamaco489/pollen-tracker/backend/internal/server"
	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

func main() {
	ctx := context.Background()
	l := logger.New()

	cfg, err := config.LoadAuthorizer()
	if err != nil {
		l.ErrorContext(ctx, "failed to load config", "error", err)
		os.Exit(1)
	}

	a := server.NewAuthServer(cfg, l)
	lambda.Start(a.Run)
}
