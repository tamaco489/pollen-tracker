package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) CreateSymptoms(ctx context.Context, req gen.CreateSymptomsRequestObject) (gen.CreateSymptomsResponseObject, error) {
	symptom, err := h.createSymptomsUseCase.CreateSymptoms(ctx, *req.Body)
	if err != nil {
		return nil, err
	}

	return gen.CreateSymptoms201JSONResponse(*symptom), nil
}
