package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	"github.com/tursodatabase/libsql-client-go/libsql"
)

const migrationsDir = "tools/migrations/sql"

func main() {
	command := flag.String("command", "up", "goose command: up, down, status, version, reset")
	flag.Parse()

	dbURL := os.Getenv("TURSO_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("TURSO_DATABASE_URL is required")
	}

	isPrd := os.Getenv("APP_ENV") == "prd"

	opts := []libsql.Option{
		libsql.WithTls(isPrd),
	}
	if token := os.Getenv("TURSO_AUTH_TOKEN"); token != "" {
		opts = append(opts, libsql.WithAuthToken(token))
	}

	connector, err := libsql.NewConnector(dbURL, opts...)
	if err != nil {
		log.Fatalf("create connector: %v", err)
	}

	db := sql.OpenDB(connector)
	defer db.Close()

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("set dialect: %v", err)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, *command, db, migrationsDir); err != nil {
		log.Fatalf("goose %s: %v", *command, err)
	}
}
