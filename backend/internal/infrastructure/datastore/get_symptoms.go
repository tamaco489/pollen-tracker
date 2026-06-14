package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/tamaco489/pollen-tracker/backend/internal/dto"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
)

// GetSymptoms はユーザーの花粉症状を日付範囲で取得する
func (r *symptomsRepository) GetSymptoms(ctx context.Context, from, to time.Time) ([]dto.GetSymptoms, error) {
	rows, err := r.queries.GetSymptoms(ctx, r.db.DB, gen.GetSymptomsParams{
		From: from.Format("2006-01-02"),
		To:   to.Format("2006-01-02"),
	})
	if err != nil {
		return nil, fmt.Errorf("list symptoms: %w", err)
	}

	result := make([]dto.GetSymptoms, 0, len(rows))
	for _, row := range rows {
		s, err := rowToSymptom(row)
		if err != nil {
			return nil, fmt.Errorf("list symptoms: %w", err)
		}
		result = append(result, *s)
	}
	return result, nil
}

func rowToSymptom(row gen.Symptom) (*dto.GetSymptoms, error) {
	id, err := uuid.Parse(row.ID)
	if err != nil {
		return nil, fmt.Errorf("parse uuid: %w", err)
	}
	date, err := time.Parse("2006-01-02", row.Date)
	if err != nil {
		return nil, fmt.Errorf("parse date: %w", err)
	}
	createdAt, err := time.Parse(time.RFC3339Nano, row.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("parse created_at: %w", err)
	}
	updatedAt, err := time.Parse(time.RFC3339Nano, row.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("parse updated_at: %w", err)
	}
	return &dto.GetSymptoms{
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
