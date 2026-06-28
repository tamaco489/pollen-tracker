package di

import (
	"context"
	"fmt"

	"github.com/tamaco489/pollen-tracker/backend/internal/handler"
	"github.com/tamaco489/pollen-tracker/backend/internal/server"
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"

	infra_datastore "github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
	google_pollen "github.com/tamaco489/pollen-tracker/backend/pkg/library/google/pollen"
)

func NewServerContainer(ctx context.Context) (*server.Server, error) {
	l := logger.New()

	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	if err := cfg.LoadSecrets(ctx); err != nil {
		return nil, fmt.Errorf("load secrets: %w", err)
	}

	conn, err := datastore.Open(ctx, cfg, l)
	if err != nil {
		return nil, fmt.Errorf("connect datastore: %w", err)
	}

	pollenClient := google_pollen.NewPollenClient(cfg.Google.PollenAPIKey)
	pollenUseCase := usecase.NewGetForecast(pollenClient)

	symptomsRepo := infra_datastore.NewSymptomsRepository(conn)
	getSymptomsUseCase := usecase.NewGetSymptoms(symptomsRepo)
	createSymptomsUseCase := usecase.NewCreateSymptoms(symptomsRepo)
	putSymptomsUseCase := usecase.NewPutSymptoms(symptomsRepo)
	getStatsUseCase := usecase.NewGetStats(symptomsRepo)
	getThresholdUseCase := usecase.NewGetThreshold(symptomsRepo)

	h := handler.NewHandler(
		pollenUseCase,
		getStatsUseCase,
		getThresholdUseCase,
		getSymptomsUseCase,
		createSymptomsUseCase,
		putSymptomsUseCase,
	)

	srv, err := server.NewServer(ctx, l, cfg, conn, h)
	if err != nil {
		return nil, fmt.Errorf("new server: %w", err)
	}

	return srv, nil
}
