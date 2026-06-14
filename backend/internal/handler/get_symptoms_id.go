package handler

import (
	"context"
	"math/rand"
	"time"

	"github.com/oapi-codegen/runtime/types"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
)

func (h *Handler) GetSymptomsId(_ context.Context, _ gen.GetSymptomsIdRequestObject) (gen.GetSymptomsIdResponseObject, error) {
	switch rand.Intn(3) {
	case 0:
		return gen.GetSymptomsId200JSONResponse{
			CreatedAt:      time.Date(2026, 5, 25, 10, 0, 0, 0, time.UTC),
			Date:           types.Date{Time: time.Date(2026, 5, 25, 0, 0, 0, 0, time.UTC)},
			Id:             types.UUID{},
			Itchy:          2,
			Note:           "stub response",
			PollenLevel:    3,
			Runny:          3,
			Sneezing:       2,
			TookMedication: true,
			UpdatedAt:      time.Date(2026, 5, 25, 10, 0, 0, 0, time.UTC),
		}, nil
	case 1:
		return gen.GetSymptomsId404JSONResponse{
			Code:  httperror.CodeNotFound.String(),
			Error: httperror.MsgNotFound.String(),
		}, nil
	default:
		return gen.GetSymptomsId500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}
}
