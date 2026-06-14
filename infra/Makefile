AWS_PROFILE ?= default
ENV ?= dev

.PHONY: fmt build lint test verify
fmt: ## コードをフォーマットする
	npx prettier --write "**/*.ts"

build: fmt ## TypeScript をコンパイルする
	npm run build

lint: ## ESLint を実行する
	npx eslint "**/*.ts"

test: ## テストを実行する
	npm test

verify: build lint test ## build, lint, test を一括実行する


.PHONY: synth diff deploy destroy
synth: build ## CloudFormation テンプレートを合成する
	ENV=$(ENV) AWS_PROFILE=$(AWS_PROFILE) npx cdk synth

diff: build ## デプロイ済みスタックとの差分を確認する
	ENV=$(ENV) AWS_PROFILE=$(AWS_PROFILE) npx cdk diff

deploy: build ## スタックをデプロイする
	ENV=$(ENV) AWS_PROFILE=$(AWS_PROFILE) npx cdk deploy --require-approval any-change

destroy: ## スタックを削除する
	ENV=$(ENV) AWS_PROFILE=$(AWS_PROFILE) npx cdk destroy


.PHONY: bootstrap
bootstrap: ## CDK bootstrap を実行する
	AWS_PROFILE=$(AWS_PROFILE) npx cdk bootstrap --qualifier $(shell node -e "const c=require('./config/env-config'); console.log(c.devConfig.bootstrapQualifier)")


.PHONY: help
help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
