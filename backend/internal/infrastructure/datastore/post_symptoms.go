package datastore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/dto"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

func (r *symptomsRepository) InsertSymptom(ctx context.Context, s *dto.CreateSymptoms) error {
	err := r.queries.InsertSymptom(ctx, r.db.DB, gen.InsertSymptomParams{
		ID:             s.ID.String(),
		Date:           s.Date.Format("2006-01-02"),
		Sneezing:       int64(s.Sneezing),
		Runny:          int64(s.Runny),
		Itchy:          int64(s.Itchy),
		PollenLevel:    int64(s.PollenLevel),
		TookMedication: int64(utils.BoolToInt(s.TookMedication)),
		Note:           s.Note,
		CreatedAt:      s.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:      s.UpdatedAt.Format(time.RFC3339Nano),
	})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("%w: symptom for this date already exists", sentinel.ErrAlreadyExists)
		}
		return fmt.Errorf("insert symptom: %w", err)
	}
	return nil
}
