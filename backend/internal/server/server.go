package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/handler"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type Server struct {
	echo *echo.Echo
}

func New(l *logger.Logger) *Server {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogMethod:   true,
			LogURI:      true,
			LogStatus:   true,
			LogLatency:  true,
			LogRemoteIP: true,
			LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
				ctx := (*c).Request().Context()
				attrs := []any{
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
					slog.String("remote_ip", v.RemoteIP),
				}
				switch {
				case v.Status >= http.StatusInternalServerError:
					l.ErrorContext(ctx, "request", attrs...)
				case v.Status >= http.StatusBadRequest:
					l.WarnContext(ctx, "request", attrs...)
				default:
					l.InfoContext(ctx, "request", attrs...)
				}
				return nil
			},
		}))

	h := handler.New(l)
	gen.RegisterHandlers(e, gen.NewStrictHandler(h, nil))

	return &Server{echo: e}
}

func (s *Server) Run(ctx context.Context) error {
	return s.echo.Start(":8080")
}
