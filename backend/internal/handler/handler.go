package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
)

type Handler struct {
	getPollenUseCase      usecase.GetForecastUseCase
	getSymptomsUseCase    usecase.GetSymptomsUseCase
	createSymptomsUseCase usecase.CreateSymptomsUseCase
}

func NewHandler(
	getPollenUseCase usecase.GetForecastUseCase,
	getSymptomsUseCase usecase.GetSymptomsUseCase,
	createSymptomsUseCase usecase.CreateSymptomsUseCase,
) *Handler {
	return &Handler{
		getPollenUseCase:      getPollenUseCase,
		getSymptomsUseCase:    getSymptomsUseCase,
		createSymptomsUseCase: createSymptomsUseCase,
	}
}
