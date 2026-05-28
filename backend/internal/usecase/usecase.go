package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain"
)

type GetForecastUseCase interface {
	GetForecast(ctx context.Context, input GetForecastInput) (*domain.PollenForecast, error)
}

type CreateSymptomsUseCase interface {
	CreateSymptoms(ctx context.Context, input CreateSymptomsInput) (*CreateSymptomsOutput, error)
}

type GetSymptomsUseCase interface {
	GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]GetSymptomsOutput, error)
}
