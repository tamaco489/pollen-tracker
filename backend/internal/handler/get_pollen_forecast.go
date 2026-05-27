package handler

import (
	"context"
	"errors"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) GetPollen(ctx context.Context, req gen.GetPollenRequestObject) (gen.GetPollenResponseObject, error) {
	input := usecase.GetForecastInput{
		Lat: req.Params.Lat,
		Lng: req.Params.Lng,
	}
	if req.Params.Date != nil {
		t := req.Params.Date.Time
		input.Date = &t
	}

	forecast, err := h.pollenUseCase.GetForecast(ctx, input)
	if err != nil {
		if errors.Is(err, sentinel.ErrInvalidInput) {
			return gen.GetPollen400JSONResponse{
				Code:  httperror.CodeBadRequest.String(),
				Error: httperror.MsgBadRequest.String(),
			}, nil
		}
		h.logger.ErrorContext(ctx, "get forecast", "error", err)
		return gen.GetPollen500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}

	return gen.GetPollen200JSONResponse{
		Date:       openapi_types.Date{Time: forecast.Date},
		Level:      int32(forecast.Level),
		PollenType: gen.GetPollen200JSONResponseBodyPollenType(forecast.PollenType),
		SeasonInfo: struct {
			Characteristics string `json:"characteristics"`
			Peak            string `json:"peak"`
			Region          string `json:"region"`
		}{
			Characteristics: forecast.SeasonInfo.Characteristics,
			Peak:            forecast.SeasonInfo.Peak,
			Region:          forecast.SeasonInfo.Region,
		},
	}, nil
}
