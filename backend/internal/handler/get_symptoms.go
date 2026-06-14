package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) GetSymptoms(ctx context.Context, req gen.GetSymptomsRequestObject) (gen.GetSymptomsResponseObject, error) {
	list, err := h.getSymptomsUseCase.GetSymptoms(ctx, req.Params)
	if err != nil {
		return nil, err
	}

	return gen.GetSymptoms200JSONResponse{
		Items: list,
		Total: int32(len(list)),
	}, nil
}
