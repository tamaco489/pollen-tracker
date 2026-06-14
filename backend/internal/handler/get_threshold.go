package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) GetThreshold(ctx context.Context, req gen.GetThresholdRequestObject) (gen.GetThresholdResponseObject, error) {
	threshold, err := h.getThresholdUseCase.GetThreshold(ctx)
	if err != nil {
		return nil, err
	}
	return gen.GetThreshold200JSONResponse(*threshold), nil
}
