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


.PHONY: build test
build: ## Lambda 関数をビルドする
	go build -o bin/bootstrap ./cmd/lambda/

test: ## テストを実行する
	go test ./...


.PHONY: up down logs
up: ## コンテナを起動する
	docker compose up -d pollen-api

down: ## コンテナを停止する
	docker compose down pollen-api

logs: ## コンテナのログを表示する
	docker compose logs -f pollen-api


.PHONY: help
help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
