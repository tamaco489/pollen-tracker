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

type PutSymptomsRepository interface {
	// UpdateSymptom は UUID でレコードを特定して症状を上書き更新する
	//
	// 対象が存在しない場合は sentinel.ErrNotFound を返す
	UpdateSymptom(ctx context.Context, s *dto.UpdateSymptoms) (*dto.GetSymptoms, error)
}
