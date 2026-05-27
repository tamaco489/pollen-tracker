package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) GetHealth(ctx context.Context, _ gen.GetHealthRequestObject) (gen.GetHealthResponseObject, error) {
	return gen.GetHealth200JSONResponse{
		Status: gen.Ok,
	}, nil
}
