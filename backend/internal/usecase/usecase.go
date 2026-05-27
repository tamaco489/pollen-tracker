package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/pollen"
)

type GetForecastUseCase interface {
	GetForecast(ctx context.Context, input GetForecastInput) (*pollen.PollenForecast, error)
}
