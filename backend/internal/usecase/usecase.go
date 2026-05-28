package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/pollen"
)

type GetForecastUseCase interface {
	GetForecast(ctx context.Context, input GetForecastInput) (*pollen.PollenForecast, error)
}

type PostSymptomsUseCase interface {
	PostSymptoms(ctx context.Context, input PostSymptomsInput) (*CreateSymptomsOutput, error)
}

type GetSymptomsUseCase interface {
	GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]GetSymptomsOutput, error)
}
