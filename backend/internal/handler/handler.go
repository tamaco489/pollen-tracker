package handler

import (
	"github.com/tamaco489/pollen-tracker/backend/internal/handler/health"
	"github.com/tamaco489/pollen-tracker/backend/internal/handler/pollen"
	"github.com/tamaco489/pollen-tracker/backend/internal/handler/symptoms"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type Handler struct {
	*health.Health
	*pollen.Pollen
	*symptoms.Symptoms
}

func New(l *logger.Logger) *Handler {
	return &Handler{
		Health:   health.New(l),
		Pollen:   pollen.New(l),
		Symptoms: symptoms.New(l),
	}
}
