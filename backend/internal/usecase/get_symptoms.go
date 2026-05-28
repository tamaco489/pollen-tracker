package usecase

import (
	"context"
	"fmt"

	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
)

type getSymptomsUseCase struct {
	repo datastore.GetSymptomsRepository
}

func NewGetSymptoms(repo datastore.GetSymptomsRepository) GetSymptomsUseCase {
	return &getSymptomsUseCase{repo: repo}
}

func (uc *getSymptomsUseCase) GetSymptoms(ctx context.Context, input GetSymptomsInput) ([]GetSymptomsOutput, error) {
	list, err := uc.repo.GetSymptoms(ctx, input.From, input.To)
	if err != nil {
		return nil, fmt.Errorf("get symptoms: %w", err)
	}

	result := make([]GetSymptomsOutput, 0, len(list))
	for _, s := range list {
		result = append(result, GetSymptomsOutput{
			ID:             s.ID,
			Date:           s.Date,
			Sneezing:       s.Sneezing,
			Runny:          s.Runny,
			Itchy:          s.Itchy,
			PollenLevel:    s.PollenLevel,
			TookMedication: s.TookMedication,
			Note:           s.Note,
			CreatedAt:      s.CreatedAt,
			UpdatedAt:      s.UpdatedAt,
		})
	}
	return result, nil
}
