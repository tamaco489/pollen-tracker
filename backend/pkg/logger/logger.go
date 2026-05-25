// Package logger provides structured logging functionality.
package logger

import (
	"context"
	"log/slog"
	"os"
)

// LogLevel represents the logging level.
type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// Config represents logger configuration.
type Config struct {
	Level  LogLevel
	Format string // "json" or "text"
}

// Logger wraps slog.Logger.
type Logger struct {
	logger *slog.Logger
}

// New creates a new Logger with default configuration.
func New() *Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return &Logger{logger: slog.New(handler)}
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}
