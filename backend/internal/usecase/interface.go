package usecase

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

type GetForecastUseCase interface {
	// GetForecast は指定座標・日付の花粉予報を取得してドメインエンティティで返す
	GetForecast(ctx context.Context, input gen.GetPollenParams) (*domain.PollenForecast, error)
}

type CreateSymptomsUseCase interface {
	// CreateSymptoms はユーザーの花粉症状を記録する
	CreateSymptoms(ctx context.Context, input gen.SymptomRequest) (*gen.SymptomResponse, error)
}

type GetSymptomsUseCase interface {
	// GetSymptoms はユーザーの花粉症状を日付範囲で取得する
	GetSymptoms(ctx context.Context, input gen.GetSymptomsParams) ([]gen.SymptomResponse, error)
}
