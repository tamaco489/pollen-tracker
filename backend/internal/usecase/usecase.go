package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain"
)

type GetForecastUseCase interface {
	GetForecast(ctx context.Context, input GetForecastInput) (*domain.PollenForecast, error)
}

type PostSymptomsUseCase interface {
	PostSymptoms(ctx context.Context, input PostSymptomsInput) (*CreateSymptomsOutput, error)
}

type GetSymptomsUseCase interface {
	GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]GetSymptomsOutput, error)
}
