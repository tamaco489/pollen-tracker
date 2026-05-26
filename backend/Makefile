.PHONY: bundle-api gen-api preview-api

OPENAPI_ROOT   := oas/oapi-base.yaml
OPENAPI_BUNDLE := oas/oapi.yaml

bundle-api: ## OpenAPI ドキュメントをバンドルする
	npx --yes @redocly/cli bundle $(OPENAPI_ROOT) -o $(OPENAPI_BUNDLE) --dereferenced

gen-api: bundle-api ## OpenAPI ドキュメントからコードを生成する
	go tool oapi-codegen --config=oas/config/types.yaml $(OPENAPI_BUNDLE)
	go tool oapi-codegen --config=oas/config/server.yaml $(OPENAPI_BUNDLE)
	sed -i '' 's|github.com/labstack/echo/v4|github.com/labstack/echo/v5|g' internal/gen/server.gen.go
	sed -i '' 's|echo\.Context|*echo.Context|g' internal/gen/server.gen.go
	sed -i '' 's|\*echo\.Route|echo.RouteInfo|g' internal/gen/server.gen.go

preview-api: ## Redoc で OpenAPI ドキュメントをブラウザプレビューする
	npx --yes @redocly/cli build-docs $(OPENAPI_ROOT) -o /tmp/pollen-tracker-api.html && open /tmp/pollen-tracker-api.html


.PHONY: fmt lint mod build test verify
fmt: ## コードをフォーマットする
	go fmt ./...

lint: ## golangci-lint を実行する
	golangci-lint run ./...

mod: ## go.mod / go.sum の整合性を確認する
	go mod tidy
	git diff --exit-code go.mod go.sum

build: mod ## Lambda 関数をビルドする
	go build -o build/bootstrap ./cmd/lambda/

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


.PHONY: db-shell
db-shell: ## Turso DB シェルに接続する
	turso db shell http://localhost:18081


.PHONY: help
help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
