package symptoms

import (
	"context"
	"math/rand"

	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/pkg/errors/httperror"
)

func (h *Symptoms) DeleteSymptomsId(_ context.Context, _ gen.DeleteSymptomsIdRequestObject) (gen.DeleteSymptomsIdResponseObject, error) {
	switch rand.Intn(3) {
	case 0:
		return gen.DeleteSymptomsId204Response{}, nil
	case 1:
		return gen.DeleteSymptomsId404JSONResponse{
			Code:  httperror.CodeNotFound.String(),
			Error: httperror.MsgNotFound.String(),
		}, nil
	default:
		return gen.DeleteSymptomsId500JSONResponse{
			Code:  httperror.CodeInternalError.String(),
			Error: httperror.MsgInternalError.String(),
		}, nil
	}
}
