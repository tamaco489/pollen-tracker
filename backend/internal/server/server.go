package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/tamaco489/pollen-tracker/backend/internal/db"
	"github.com/tamaco489/pollen-tracker/backend/internal/gen"
	"github.com/tamaco489/pollen-tracker/backend/internal/handler"
	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *db.DB
	cancel atomic.Pointer[context.CancelFunc]
	done   chan struct{}
}

// New は設定・DB 接続・ルーターを初期化して Server を返す。
func New(ctx context.Context, l *logger.Logger) (*Server, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	conn, err := db.Open(ctx, cfg, l)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	e := echo.New()

	// Recover を最初に追加して panic をサーバークラッシュではなく 500 に変換する。
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogMethod:   true,
			LogURI:      true,
			LogStatus:   true,
			LogLatency:  true,
			LogRemoteIP: true,
			LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
				reqCtx := (*c).Request().Context()
				attrs := []any{
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.Duration("latency", v.Latency),
					slog.String("remote_ip", v.RemoteIP),
				}
				switch {
				case v.Status >= http.StatusInternalServerError:
					l.ErrorContext(reqCtx, "request", attrs...)
				case v.Status >= http.StatusBadRequest:
					l.WarnContext(reqCtx, "request", attrs...)
				default:
					l.InfoContext(reqCtx, "request", attrs...)
				}
				return nil
			},
		}))

	h := handler.New(l)
	gen.RegisterHandlers(e, gen.NewStrictHandler(h, nil))

	l.InfoContext(ctx, "server initialized",
		"addr", ":"+cfg.App.Port,
		"env", cfg.App.Env,
		"service", cfg.App.ServiceName(),
	)

	return &Server{
		echo: e,
		cfg:  cfg,
		db:   conn,
		done: make(chan struct{}),
	}, nil
}

// Run は HTTP サーバーを起動してブロックする。
// Shutdown が呼ばれると graceful shutdown が完了してから返る。
func (s *Server) Run(ctx context.Context) error {
	runCtx, cancelFn := context.WithCancel(ctx)
	s.cancel.Store(&cancelFn)
	defer close(s.done)

	isPrd := s.cfg.App.Env.IsProduction()
	sc := echo.StartConfig{
		Address:         ":" + s.cfg.App.Port,
		GracefulTimeout: 10 * time.Second,
		// 本番環境ではバナー・ポート出力を抑制して構造化ログに統一する。
		// 開発環境では Echo のアスキーアートをそのまま表示する。
		HideBanner: isPrd,
		HidePort:   isPrd,
	}
	return sc.Start(runCtx, s.echo)
}

// Shutdown は HTTP サーバーの graceful shutdown と DB 接続のクローズを行う。
func (s *Server) Shutdown(ctx context.Context) error {
	if fn := s.cancel.Load(); fn != nil {
		(*fn)()
	}
	// HTTP サーバーの graceful shutdown が完了するまで待機する。
	select {
	case <-s.done:
	case <-ctx.Done():
	}
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("db close: %w", err)
	}
	return nil
}
