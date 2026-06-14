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

type putSymptomsUseCase struct {
	repo datastore.PutSymptomsRepository
}

func NewPutSymptoms(repo datastore.PutSymptomsRepository) PutSymptomsUseCase {
	return &putSymptomsUseCase{repo: repo}
}

// PutSymptoms は UUID で症状ログを上書き更新する
func (uc *putSymptomsUseCase) PutSymptoms(ctx context.Context, id uuid.UUID, input gen.SymptomUpdateRequest) (*gen.SymptomResponse, error) {
	if err := uc.validate(input); err != nil {
		return nil, err
	}

	s := &dto.UpdateSymptoms{
		ID:             id,
		Sneezing:       input.Sneezing,
		Runny:          input.Runny,
		Itchy:          input.Itchy,
		PollenLevel:    input.PollenLevel,
		TookMedication: input.TookMedication,
		Note:           input.Note,
		UpdatedAt:      time.Now().UTC(),
	}

	updated, err := uc.repo.UpdateSymptom(ctx, s)
	if err != nil {
		return nil, err
	}

	return &gen.SymptomResponse{
		Id:             types.UUID(updated.ID),
		Date:           types.Date{Time: updated.Date},
		Sneezing:       updated.Sneezing,
		Runny:          updated.Runny,
		Itchy:          updated.Itchy,
		PollenLevel:    updated.PollenLevel,
		TookMedication: updated.TookMedication,
		Note:           updated.Note,
		CreatedAt:      updated.CreatedAt,
		UpdatedAt:      updated.UpdatedAt,
	}, nil
}

func (uc *putSymptomsUseCase) validate(input gen.SymptomUpdateRequest) error {
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
