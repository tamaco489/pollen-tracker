package datastore

import (
	"context"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
)

type GetSymptomsRepository interface {
	GetSymptoms(ctx context.Context, from, to time.Time) ([]symptoms.Symptom, error)
}

type PostSymptomsRepository interface {
	Insert(ctx context.Context, s *symptoms.Symptom) error
}
