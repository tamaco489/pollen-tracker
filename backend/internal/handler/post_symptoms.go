package handler

import (
	"context"
	"math/rand"
	"time"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) PostSymptoms(_ context.Context, _ gen.PostSymptomsRequestObject) (gen.PostSymptomsResponseObject, error) {
	switch rand.Intn(4) {
	case 0:
		return gen.PostSymptoms201JSONResponse{
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
		}, nil
	case 1:
		return gen.PostSymptoms400JSONResponse{
			Code:  httperror.CodeBadRequest.String(),
			Error: httperror.MsgBadRequest.String(),
		}, nil
	case 2:
		return gen.PostSymptoms409JSONResponse{
			Code:  httperror.CodeAlreadyExists.String(),
			Error: httperror.MsgAlreadyExists.String(),
		}, nil
	default:
		return gen.PostSymptoms500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}
}
