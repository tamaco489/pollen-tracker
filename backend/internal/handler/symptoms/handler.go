package symptoms

import "github.com/tamaco489/pollen-tracker/backend/pkg/logger"

type Symptoms struct {
	logger *logger.Logger
}

func New(l *logger.Logger) *Symptoms { return &Symptoms{logger: l} }
