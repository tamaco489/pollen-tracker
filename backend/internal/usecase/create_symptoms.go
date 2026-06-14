package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"

	"github.com/tamaco489/pollen-tracker/backend/internal/dto"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
)

type createSymptomsUseCase struct {
	repo datastore.CreateSymptomsRepository
}

func NewCreateSymptoms(repo datastore.CreateSymptomsRepository) CreateSymptomsUseCase {
	return &createSymptomsUseCase{repo: repo}
}

const (
	symptomLevelMin = 1
	symptomLevelMax = 5
	noteMaxChars    = 200
)

// CreateSymptoms はユーザーの花粉症状を記録する
func (uc *createSymptomsUseCase) CreateSymptoms(ctx context.Context, input gen.SymptomRequest) (*gen.SymptomResponse, error) {
	if err := uc.validate(input); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("generate uuid: %w", err)
	}

	now := time.Now().UTC()
	s := &dto.CreateSymptoms{
		ID:             id,
		Date:           input.Date.Time.UTC().Truncate(24 * time.Hour),
		Sneezing:       input.Sneezing,
		Runny:          input.Runny,
		Itchy:          input.Itchy,
		PollenLevel:    input.PollenLevel,
		TookMedication: input.TookMedication,
		Note:           input.Note,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := uc.repo.InsertSymptom(ctx, s); err != nil {
		return nil, err
	}

	return &gen.SymptomResponse{
		Id:             types.UUID(s.ID),
		Date:           types.Date{Time: s.Date},
		Sneezing:       s.Sneezing,
		Runny:          s.Runny,
		Itchy:          s.Itchy,
		PollenLevel:    s.PollenLevel,
		TookMedication: s.TookMedication,
		Note:           s.Note,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}, nil
}

func (uc *createSymptomsUseCase) validate(input gen.SymptomRequest) error {
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
