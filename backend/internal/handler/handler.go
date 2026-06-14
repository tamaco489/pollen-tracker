package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
)

type Handler struct {
	getPollenUseCase      usecase.GetForecastUseCase
	getSymptomsUseCase    usecase.GetSymptomsUseCase
	createSymptomsUseCase usecase.CreateSymptomsUseCase
	putSymptomsUseCase    usecase.PutSymptomsUseCase
}

func NewHandler(
	getPollenUseCase usecase.GetForecastUseCase,
	getSymptomsUseCase usecase.GetSymptomsUseCase,
	createSymptomsUseCase usecase.CreateSymptomsUseCase,
	putSymptomsUseCase usecase.PutSymptomsUseCase,
) *Handler {
	return &Handler{
		getPollenUseCase:      getPollenUseCase,
		getSymptomsUseCase:    getSymptomsUseCase,
		createSymptomsUseCase: createSymptomsUseCase,
		putSymptomsUseCase:    putSymptomsUseCase,
	}
}
