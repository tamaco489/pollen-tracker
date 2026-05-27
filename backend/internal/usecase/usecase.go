package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/pkg/library/google/pollen"
)

// GetForecastUseCase は Google Pollen API 呼び出しの依存インターフェース
type GetForecastUseCase interface {
	GetForecast(ctx context.Context, req *pollen.ForecastRequest) (*pollen.ForecastResponse, error)
}
