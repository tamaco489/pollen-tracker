package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/pollen"
	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
)

type GetForecastUseCase interface {
	GetForecast(ctx context.Context, input GetForecastInput) (*pollen.PollenForecast, error)
}

type PostSymptomsUseCase interface {
	PostSymptoms(ctx context.Context, input PostSymptomsInput) (*symptoms.Symptom, error)
}

type GetSymptomsUseCase interface {
	GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]symptoms.Symptom, error)
}
