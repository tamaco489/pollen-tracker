package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) GetStats(ctx context.Context, req gen.GetStatsRequestObject) (gen.GetStatsResponseObject, error) {
	stats, err := h.getStatsUseCase.GetStats(ctx, req.Params)
	if err != nil {
		return nil, err
	}

	return gen.GetStats200JSONResponse(*stats), nil
}
