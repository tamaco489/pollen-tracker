.PHONY: setup-tools
setup-tools: ## 開発環境をセットアップする
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/tursodatabase/turso-cli/cmd/turso@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	npm install -g @redocly/cli


.PHONY: bundle-api gen-api preview-api

OPENAPI_ROOT   := oas/oapi-base.yaml
OPENAPI_BUNDLE := oas/oapi.yaml

bundle-api: ## OpenAPI ドキュメントをバンドルする
	npx --yes @redocly/cli bundle $(OPENAPI_ROOT) -o $(OPENAPI_BUNDLE)

gen-api: bundle-api ## OpenAPI ドキュメントからコードを生成する
	go tool oapi-codegen --config=oas/config/types.yaml $(OPENAPI_BUNDLE)
	go tool oapi-codegen --config=oas/config/server.yaml $(OPENAPI_BUNDLE)
	sed -i '' 's|github.com/labstack/echo/v4|github.com/labstack/echo/v5|g' internal/gen/server.gen.go
	sed -i '' 's|echo\.Context|*echo.Context|g' internal/gen/server.gen.go
	sed -i '' 's|\*echo\.Route|echo.RouteInfo|g' internal/gen/server.gen.go

preview-api: ## Redoc で OpenAPI ドキュメントをブラウザプレビューする
	npx --yes @redocly/cli build-docs $(OPENAPI_ROOT) -o /tmp/pollen-tracker-api.html && open /tmp/pollen-tracker-api.html


.PHONY: gen-sqlc sqlc-diff sqlc-vet sqlc-lint
gen-sqlc: ## sqlc からコードを生成する
	sqlc generate -f tools/sqlc/sqlc.yaml

sqlc-diff: ## sqlc スキーマの差分を確認する
	sqlc diff -f tools/sqlc/sqlc.yaml

sqlc-vet: ## sqlc 静的解析を実行する
	sqlc vet -f tools/sqlc/sqlc.yaml

sqlc-lint: sqlc-diff sqlc-vet ## sqlc スキーマ差分確認と静的解析を実行する


.PHONY: fmt lint lint-api lint-pkg lint-tools mod build test verify
fmt: ## コードをフォーマットする
	go fmt ./...

lint: ## 全ディレクトリを lint する
	golangci-lint run ./...

lint-api: ## API・Authorizer Lambda を lint する
	golangci-lint run ./cmd/api/... ./cmd/authorizer/... ./internal/...

lint-pkg: ## 共通パッケージを lint する
	golangci-lint run ./pkg/...

lint-tools: ## tools を lint する
	golangci-lint run ./tools/...

mod: ## go.mod / go.sum の整合性を確認する
	go mod tidy
	git diff --exit-code go.mod go.sum

build: mod ## API・Authorizer Lambda 関数をまとめてビルドする
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o build/api/bootstrap ./cmd/api/
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o build/authorizer/bootstrap ./cmd/authorizer/


test: ## テストを実行する
	go test ./...

verify: fmt lint build test ## fmt, lint, mod, build, test を一括実行する


.PHONY: up down logs
up: ## コンテナを起動する
	docker compose up -d pollen-db pollen-api

down: ## コンテナを停止する
	docker compose down pollen-db pollen-api

logs: ## コンテナのログを表示する
	docker compose logs -f pollen-api


.PHONY: migrate-create migrate-up migrate-down migrate-status db-shell
migrate-create: ## マイグレーションファイルを作成する (name=xxx が必須)
	goose -dir tools/migrations/sql create $(name) sql

migrate-up: ## DB マイグレーションを実行する
	docker compose exec pollen-api go run ./tools/migrations/cmd/ -command up

migrate-down: ## DB マイグレーションを1つロールバックする
	docker compose exec pollen-api go run ./tools/migrations/cmd/ -command down

migrate-status: ## マイグレーションの適用状況を確認する
	docker compose exec pollen-api go run ./tools/migrations/cmd/ -command status

db-shell: ## Turso DB シェルに接続する
	turso db shell http://localhost:18081


ENV              ?= dev
ARTIFACTS_BUCKET ?= $(ENV)-pollen-tracker-storage

.PHONY: upload-api upload-authorizer
upload-api: ## API Lambda バイナリを zip 化して S3 にアップロードする
	zip -j build/api/bootstrap.zip build/api/bootstrap
	aws s3 cp build/api/bootstrap.zip s3://$(ARTIFACTS_BUCKET)/artifacts/api/bootstrap.zip

upload-authorizer: ## Authorizer Lambda バイナリを zip 化して S3 にアップロードする
	zip -j build/authorizer/bootstrap.zip build/authorizer/bootstrap
	aws s3 cp build/authorizer/bootstrap.zip s3://$(ARTIFACTS_BUCKET)/artifacts/authorizer/bootstrap.zip

deploy-api: upload-api ## API Lambda を S3 にアップロードしてデプロイする
	aws lambda update-function-code \
		--function-name $(ENV)-pollen-tracker-api \
		--s3-bucket $(ARTIFACTS_BUCKET) \
		--s3-key artifacts/api/bootstrap.zip \
		--profile $(AWS_PROFILE) | jq .

deploy-authorizer: upload-authorizer ## Authorizer Lambda を S3 にアップロードしてデプロイする
	aws lambda update-function-code \
		--function-name $(ENV)-pollen-tracker-authorizer \
		--s3-bucket $(ARTIFACTS_BUCKET) \
		--s3-key artifacts/authorizer/bootstrap.zip \
		--profile $(AWS_PROFILE) | jq .


.PHONY: help
help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
