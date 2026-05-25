package symptoms

import (
	"context"
	"math/rand"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Symptoms) GetSymptoms(_ context.Context, _ gen.GetSymptomsRequestObject) (gen.GetSymptomsResponseObject, error) {
	switch rand.Intn(3) {
	case 0:
		return gen.GetSymptoms200JSONResponse{
			Items: []struct {
				CreatedAt      time.Time          `json:"created_at"`
				Date           openapi_types.Date `json:"date"`
				Id             openapi_types.UUID `json:"id"`
				Itchy          int32              `json:"itchy"`
				Note           string             `json:"note"`
				PollenLevel    int32              `json:"pollen_level"`
				Runny          int32              `json:"runny"`
				Sneezing       int32              `json:"sneezing"`
				TookMedication bool               `json:"took_medication"`
				UpdatedAt      time.Time          `json:"updated_at"`
			}{
				{
					CreatedAt:      time.Date(2026, 5, 25, 10, 0, 0, 0, time.UTC),
					Date:           openapi_types.Date{Time: time.Date(2026, 5, 25, 0, 0, 0, 0, time.UTC)},
					Id:             openapi_types.UUID{},
					Itchy:          2,
					Note:           "stub response",
					PollenLevel:    3,
					Runny:          3,
					Sneezing:       2,
					TookMedication: true,
					UpdatedAt:      time.Date(2026, 5, 25, 10, 0, 0, 0, time.UTC),
				},
			},
			Total: 1,
		}, nil
	case 1:
		return gen.GetSymptoms400JSONResponse{
			Code:  httperror.CodeBadRequest.String(),
			Error: httperror.MsgBadRequest.String(),
		}, nil
	default:
		return gen.GetSymptoms500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}
}
