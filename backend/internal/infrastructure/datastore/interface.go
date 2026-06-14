package datastore

import (
	"context"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/dto"
)

type GetSymptomsRepository interface {
	// GetSymptoms はユーザーの花粉症状を日付範囲で取得する
	GetSymptoms(ctx context.Context, from, to time.Time) ([]dto.GetSymptoms, error)
}

type CreateSymptomsRepository interface {
	// InsertSymptom はユーザーの花粉症状を記録する
	InsertSymptom(ctx context.Context, s *dto.CreateSymptoms) error
}
