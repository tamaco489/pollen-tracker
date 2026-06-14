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

type GetStatsRepository interface {
	// GetWeeklyStats は週次集計データを取得する
	GetWeeklyStats(ctx context.Context, from, to time.Time) ([]dto.StatsItem, error)

	// GetMonthlyStats は月次集計データを取得する
	GetMonthlyStats(ctx context.Context, from, to time.Time) ([]dto.StatsItem, error)
}

type GetThresholdRepository interface {
	// GetSymptomPollenLevels は症状が記録された日の花粉レベル一覧を昇順で返す
	GetSymptomPollenLevels(ctx context.Context) ([]int64, error)
}
