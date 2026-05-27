package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
)

type Handler struct {
	pollenUseCase usecase.GetForecastUseCase
}

func New(pollenUseCase usecase.GetForecastUseCase) *Handler {
	return &Handler{
		pollenUseCase: pollenUseCase,
	}
}
