package health

import "github.com/tamaco489/pollen-tracker/backend/pkg/logger"

type Health struct {
	logger *logger.Logger
}

func New(l *logger.Logger) *Health { return &Health{logger: l} }
