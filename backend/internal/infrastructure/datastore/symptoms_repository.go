package datastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

type symptomsRepository struct {
	db *datastore.DB
}

func NewSymptomsRepository(db *datastore.DB) Repository {
	return &symptomsRepository{db: db}
}

func (r *symptomsRepository) Insert(ctx context.Context, s *symptoms.Symptom) error {
	const q = `
INSERT INTO symptoms (
	id, date, sneezing, runny, itchy,
	pollen_level, took_medication, note,
	created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, q,
		s.ID.String(),
		s.Date.Format("2006-01-02"),
		s.Sneezing,
		s.Runny,
		s.Itchy,
		s.PollenLevel,
		utils.BoolToInt(s.TookMedication),
		s.Note,
		s.CreatedAt.Format(time.RFC3339Nano),
		s.UpdatedAt.Format(time.RFC3339Nano),
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("%w: symptom for this date already exists", sentinel.ErrAlreadyExists)
		}
		return fmt.Errorf("insert symptom: %w", err)
	}
	return nil
}
