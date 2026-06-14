package handler

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"
)

func (h *Handler) PutSymptomsId(ctx context.Context, req gen.PutSymptomsIdRequestObject) (gen.PutSymptomsIdResponseObject, error) {
	id := uuid.UUID(req.Id)

	symptom, err := h.putSymptomsUseCase.PutSymptoms(ctx, id, *req.Body)
	if err != nil {
		switch {
		case errors.Is(err, sentinel.ErrInvalidInput):
			return gen.PutSymptomsId400JSONResponse{
				Code:  httperror.CodeBadRequest.String(),
				Error: httperror.MsgBadRequest.String(),
			}, nil
		case errors.Is(err, sentinel.ErrNotFound):
			return gen.PutSymptomsId404JSONResponse{
				Code:  httperror.CodeNotFound.String(),
				Error: httperror.MsgNotFound.String(),
			}, nil
		default:
			return gen.PutSymptomsId500JSONResponse{
				Code:  httperror.CodeInternalError.String(),
				Error: httperror.MsgInternalError.String(),
			}, nil
		}
	}

	return gen.PutSymptomsId200JSONResponse{
		Id:             symptom.Id,
		Date:           symptom.Date,
		Sneezing:       symptom.Sneezing,
		Runny:          symptom.Runny,
		Itchy:          symptom.Itchy,
		PollenLevel:    symptom.PollenLevel,
		TookMedication: symptom.TookMedication,
		Note:           symptom.Note,
		CreatedAt:      symptom.CreatedAt,
		UpdatedAt:      symptom.UpdatedAt,
	}, nil
}
