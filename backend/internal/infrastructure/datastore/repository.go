package datastore

import (
	"context"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/dto"
)

type GetSymptomsRepository interface {
	GetSymptoms(ctx context.Context, from, to time.Time) ([]dto.GetSymptoms, error)
}

type PostSymptomsRepository interface {
	InsertSymptom(ctx context.Context, s *dto.CreateSymptoms) error
}
