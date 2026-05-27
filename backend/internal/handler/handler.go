package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/usecase"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type Handler struct {
	logger        *logger.Logger
	pollenUseCase usecase.GetForecastUseCase
}

func New(l *logger.Logger, pollenUseCase usecase.GetForecastUseCase) *Handler {
	return &Handler{
		logger:        l,
		pollenUseCase: pollenUseCase,
	}
}
