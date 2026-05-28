package datastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

type symptomsRepository struct {
	db      *datastore.DB
	queries *gen.Queries
}

func NewSymptomsRepository(db *datastore.DB) *symptomsRepository {
	return &symptomsRepository{db: db, queries: gen.New()}
}

func (r *symptomsRepository) GetSymptoms(ctx context.Context, from, to time.Time) ([]symptoms.Symptom, error) {
	rows, err := r.queries.GetSymptoms(ctx, r.db.DB, gen.GetSymptomsParams{
		From: from.Format("2006-01-02"),
		To:   to.Format("2006-01-02"),
	})
	if err != nil {
		return nil, fmt.Errorf("list symptoms: %w", err)
	}

	result := make([]symptoms.Symptom, 0, len(rows))
	for _, row := range rows {
		s, err := rowToSymptom(row)
		if err != nil {
			return nil, fmt.Errorf("list symptoms: %w", err)
		}
		result = append(result, s)
	}
	return result, nil
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

func rowToSymptom(row gen.Symptom) (symptoms.Symptom, error) {
	id, err := uuid.Parse(row.ID)
	if err != nil {
		return symptoms.Symptom{}, fmt.Errorf("parse uuid: %w", err)
	}
	date, err := time.Parse("2006-01-02", row.Date)
	if err != nil {
		return symptoms.Symptom{}, fmt.Errorf("parse date: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339Nano, row.CreatedAt)
	if err != nil {
		return symptoms.Symptom{}, fmt.Errorf("parse created_at: %w", err)
	}
	updatedAt, err := time.Parse(time.RFC3339Nano, row.UpdatedAt)
	if err != nil {
		return symptoms.Symptom{}, fmt.Errorf("parse updated_at: %w", err)
	}
	return symptoms.Symptom{
		ID:             id,
		Date:           date,
		Sneezing:       int32(row.Sneezing),
		Runny:          int32(row.Runny),
		Itchy:          int32(row.Itchy),
		PollenLevel:    int32(row.PollenLevel),
		TookMedication: row.TookMedication != 0,
		Note:           row.Note,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}
