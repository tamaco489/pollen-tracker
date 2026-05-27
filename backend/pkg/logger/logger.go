// Package logger は構造化ログ機能を提供する
package logger

import (
	"context"
	"log/slog"
	"os"
)

// LogLevel はログレベルを表す型
type LogLevel string

const (
	LevelDebug LogLevel = "debug"
	LevelInfo  LogLevel = "info"
	LevelWarn  LogLevel = "warn"
	LevelError LogLevel = "error"
)

// Config はロガーの設定
type Config struct {
	Level  LogLevel
	Format string // "json" or "text"
}

// Logger は slog.Logger のラッパー
type Logger struct {
	logger *slog.Logger
}

// New はデフォルト設定で Logger を生成して返す
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
