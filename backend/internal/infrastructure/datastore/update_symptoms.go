package datastore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/dto"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

func (r *symptomsRepository) UpdateSymptom(ctx context.Context, s *dto.UpdateSymptoms) (*dto.GetSymptoms, error) {
	row, err := r.queries.UpdateSymptom(ctx, r.db.DB, gen.UpdateSymptomParams{
		ID:             s.ID.String(),
		Sneezing:       int64(s.Sneezing),
		Runny:          int64(s.Runny),
		Itchy:          int64(s.Itchy),
		PollenLevel:    int64(s.PollenLevel),
		TookMedication: int64(utils.BoolToInt(s.TookMedication)),
		Note:           s.Note,
		UpdatedAt:      s.UpdatedAt.Format(time.RFC3339Nano),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: symptom not found", sentinel.ErrNotFound)
		}
		return nil, fmt.Errorf("update symptom: %w", err)
	}

	result, err := rowToSymptom(row)
	if err != nil {
		return nil, fmt.Errorf("update symptom: %w", err)
	}
	return result, nil
}
