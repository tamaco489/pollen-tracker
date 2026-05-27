package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
	"github.com/tamaco489/pollen-tracker/backend/pkg/logger"
)

// DB は *sql.DB のラッパーOpen / Close を同一パッケージで管理する
type DB struct {
	*sql.DB
}

// Open は config から Turso DB への接続を初期化して返す
func Open(ctx context.Context, cfg *config.Config, l *logger.Logger) (*DB, error) {
	opts := []libsql.Option{
		libsql.WithTls(cfg.App.Env.IsProduction()),
	}

	if cfg.TursoDB.AuthToken != "" {
		opts = append(opts, libsql.WithAuthToken(cfg.TursoDB.AuthToken))
	}

	connector, err := libsql.NewConnector(cfg.TursoDB.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("create db connector: %w", err)
	}

	db := sql.OpenDB(connector)

	// Lambda 向けコネクションプール設定: 1インスタンスにつき接続1本で十分
	// MaxIdleConns=1 でウォームインスタンス間の接続を再利用し、30s アイドル後に解放してコールドスタート時のリソースを節約する
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(30 * time.Second)

	u, _ := url.Parse(cfg.TursoDB.URL)
	l.InfoContext(ctx, "database connection initialized",
		"host", u.Host,
		"max_open_conns", 1,
		"tls", cfg.App.Env.IsProduction(),
	)

	return &DB{db}, nil
}
