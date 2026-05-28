package datastore

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
)

type Repository interface {
	Insert(ctx context.Context, s *symptoms.Symptom) error
}
