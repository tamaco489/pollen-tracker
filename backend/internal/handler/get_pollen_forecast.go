package handler

import (
	"context"

	"github.com/oapi-codegen/runtime/types"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
)

func (h *Handler) GetPollen(ctx context.Context, req gen.GetPollenRequestObject) (gen.GetPollenResponseObject, error) {
	forecast, err := h.getPollenUseCase.GetForecast(ctx, req.Params)
	if err != nil {
		return nil, err
	}

	return gen.GetPollen200JSONResponse{
		Date:       types.Date{Time: forecast.Date},
		Level:      int32(forecast.Level),
		PollenType: gen.PollenResponsePollenType(forecast.PollenType),
		SeasonInfo: gen.SeasonInfo{
			Characteristics: forecast.SeasonInfo.Characteristics,
			Peak:            forecast.SeasonInfo.Peak,
			Region:          forecast.SeasonInfo.Region,
		},
	}, nil
}
