package pollen

import "github.com/tamaco489/pollen-tracker/backend/pkg/logger"

type Pollen struct {
	logger *logger.Logger
}

func New(l *logger.Logger) *Pollen { return &Pollen{logger: l} }
