package pollen

import (
	"context"
	"math/rand"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Pollen) GetPollen(_ context.Context, _ gen.GetPollenRequestObject) (gen.GetPollenResponseObject, error) {
	switch rand.Intn(3) {
	case 0:
		return gen.GetPollen200JSONResponse{
			Date:       openapi_types.Date{Time: time.Date(2026, 5, 25, 0, 0, 0, 0, time.UTC)},
			Level:      3,
			PollenType: gen.CEDAR,
			SeasonInfo: struct {
				Characteristics string `json:"characteristics"`
				Peak            string `json:"peak"`
				Region          string `json:"region"`
			}{
				Characteristics: "目のかゆみ・鼻水が多い",
				Peak:            "2月〜4月",
				Region:          "関東・近畿",
			},
		}, nil
	case 1:
		return gen.GetPollen400JSONResponse{
			Code:  httperror.CodeBadRequest.String(),
			Error: httperror.MsgBadRequest.String(),
		}, nil
	default:
		return gen.GetPollen500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}
}
