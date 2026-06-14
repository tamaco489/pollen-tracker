package usecase

import (
	"context"
	"fmt"

	"github.com/oapi-codegen/runtime/types"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/infrastructure/datastore"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"
)

type getSymptomsUseCase struct {
	repo datastore.GetSymptomsRepository
}

func NewGetSymptoms(repo datastore.GetSymptomsRepository) GetSymptomsUseCase {
	return &getSymptomsUseCase{repo: repo}
}

// GetSymptoms はユーザーの花粉症状を日付範囲で取得する
func (uc *getSymptomsUseCase) GetSymptoms(ctx context.Context, input gen.GetSymptomsParams) ([]gen.SymptomResponse, error) {
	from, to := utils.DefaultDateRange()
	if input.From != nil {
		from = input.From.Time
	}
	if input.To != nil {
		to = input.To.Time
	}

	list, err := uc.repo.GetSymptoms(ctx, from, to)
	if err != nil {
		return nil, fmt.Errorf("get symptoms: %w", err)
	}

	result := make([]gen.SymptomResponse, 0, len(list))
	for _, s := range list {
		result = append(result, gen.SymptomResponse{
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
		})
	}

	return result, nil
}
