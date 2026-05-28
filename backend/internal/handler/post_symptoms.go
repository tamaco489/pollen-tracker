package handler

import (
	"context"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) PostSymptoms(ctx context.Context, req gen.PostSymptomsRequestObject) (gen.PostSymptomsResponseObject, error) {
	input := usecase.PostSymptomsInput{
		Date:           req.Body.Date.Time,
		Sneezing:       req.Body.Sneezing,
		Runny:          req.Body.Runny,
		Itchy:          req.Body.Itchy,
		PollenLevel:    req.Body.PollenLevel,
		TookMedication: req.Body.TookMedication,
		Note:           req.Body.Note,
	}

	symptom, err := h.postSymptomsUseCase.PostSymptoms(ctx, input)
	if err != nil {
		return nil, err
	}

	return gen.PostSymptoms201JSONResponse{
		Id:             openapi_types.UUID(symptom.ID),
		Date:           openapi_types.Date{Time: symptom.Date},
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
