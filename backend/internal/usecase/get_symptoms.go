package usecase

import (
	"context"
	"fmt"

	"github.com/tamaco489/pollen-tracker/backend/internal/domain/symptoms"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
)

type getSymptomsUseCase struct {
	repo datastore.GetSymptomsRepository
}

func NewGetSymptoms(repo datastore.GetSymptomsRepository) GetSymptomsUseCase {
	return &getSymptomsUseCase{repo: repo}
}

func (uc *getSymptomsUseCase) GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]symptoms.Symptom, error) {
	list, err := uc.repo.GetSymptoms(ctx, input.From, input.To)
	if err != nil {
		return nil, fmt.Errorf("get symptoms: %w", err)
	}
	return list, nil
}
