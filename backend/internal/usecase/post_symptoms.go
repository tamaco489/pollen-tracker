package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
)

type postSymptomsUseCase struct {
	repo datastore.PostSymptomsRepository
}

func NewPostSymptoms(repo datastore.PostSymptomsRepository) PostSymptomsUseCase {
	return &postSymptomsUseCase{repo: repo}
}

const (
	symptomLevelMin = 1
	symptomLevelMax = 5
	noteMaxChars    = 200
)

func (uc *postSymptomsUseCase) PostSymptoms(ctx context.Context, input PostSymptomsInput) (*symptoms.Symptom, error) {
	if err := uc.validate(input); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("generate uuid: %w", err)
	}

	now := time.Now().UTC()
	s := &symptoms.Symptom{
		ID:             id,
		Date:           input.Date.UTC().Truncate(24 * time.Hour),
		Sneezing:       input.Sneezing,
		Runny:          input.Runny,
		Itchy:          input.Itchy,
		PollenLevel:    input.PollenLevel,
		TookMedication: input.TookMedication,
		Note:           input.Note,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := uc.repo.Insert(ctx, s); err != nil {
		return nil, err
	}

	return s, nil
}

func (uc *postSymptomsUseCase) validate(input PostSymptomsInput) error {
	for _, v := range []struct {
		val  int32
		name string
	}{
		{input.Sneezing, "sneezing"},
		{input.Runny, "runny"},
		{input.Itchy, "itchy"},
		{input.PollenLevel, "pollen_level"},
	} {
		if v.val < symptomLevelMin || v.val > symptomLevelMax {
			return fmt.Errorf(
				"%w: %s must be between %d and %d",
				sentinel.ErrInvalidInput,
				v.name,
				symptomLevelMin,
				symptomLevelMax,
			)
		}
	}

	if len([]rune(input.Note)) > noteMaxChars {
		return fmt.Errorf("%w: note must not exceed %d characters", sentinel.ErrInvalidInput, noteMaxChars)
	}

	return nil
}
