package logger

import (
	"log/slog"
	"os"
)

var defaultLogger *Logger

func Init(config Config) {
	var level slog.Level
	switch config.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	}

	var handler slog.Handler
	if config.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	defaultLogger = &Logger{logger: slog.New(handler)}
	slog.SetDefault(defaultLogger.logger)
}

func InitDefault() {
	Init(Config{Level: LevelInfo, Format: "json"})
}

func GetLogger() *Logger {
	if defaultLogger == nil {
		InitDefault()
	}
	return defaultLogger
}
