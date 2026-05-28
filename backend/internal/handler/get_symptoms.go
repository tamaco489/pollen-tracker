package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
	"github.com/tamaco489/pollen-tracker/backend/pkg/utils"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) GetSymptoms(ctx context.Context, req gen.GetSymptomsRequestObject) (gen.GetSymptomsResponseObject, error) {
	from, to := utils.DefaultDateRange()
	if req.Params.From != nil {
		from = req.Params.From.Time
	}
	if req.Params.To != nil {
		to = req.Params.To.Time
	}

	list, err := h.getSymptomsUseCase.GetSymptoms(ctx, usecase.GetSymptomsInput{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	items := make([]gen.SymptomResponse, 0, len(list))
	for _, s := range list {
		items = append(items, gen.SymptomResponse{
			Id:             openapi_types.UUID(s.ID),
			Date:           openapi_types.Date{Time: s.Date},
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

	return gen.GetSymptoms200JSONResponse{
		Items: items,
		Total: int32(len(items)),
	}, nil
}
