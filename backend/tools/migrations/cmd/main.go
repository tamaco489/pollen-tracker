package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/tamaco489/pollen-tracker/backend/pkg/config"
)

const dir = "tools/migrations/sql"

func main() {
	command := flag.String("command", "up", "goose command: up, down, status, version, reset")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	opts := []libsql.Option{
		libsql.WithTls(cfg.App.Env.IsProduction()),
	}
	if cfg.TursoDB.AuthToken != "" {
		opts = append(opts, libsql.WithAuthToken(cfg.TursoDB.AuthToken))
	}

	connector, err := libsql.NewConnector(cfg.TursoDB.URL, opts...)
	if err != nil {
		log.Fatalf("create connector: %v", err)
	}

	db := sql.OpenDB(connector)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("close db: %v", err)
		}
	}()

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, *command, db, dir); err != nil {
		log.Fatalf("goose %s: %v", *command, err)
	}
}
