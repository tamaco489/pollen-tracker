package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
)

type Handler struct {
	getPollenUseCase    usecase.GetForecastUseCase
	getSymptomsUseCase  usecase.GetSymptomsUseCase
	postSymptomsUseCase usecase.PostSymptomsUseCase
}

func NewHandler(
	getPollenUseCase usecase.GetForecastUseCase,
	getSymptomsUseCase usecase.GetSymptomsUseCase,
	postSymptomsUseCase usecase.PostSymptomsUseCase,
) *Handler {
	return &Handler{
		getPollenUseCase:    getPollenUseCase,
		getSymptomsUseCase:  getSymptomsUseCase,
		postSymptomsUseCase: postSymptomsUseCase,
	}
}
