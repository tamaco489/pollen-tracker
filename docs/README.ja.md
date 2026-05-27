# pollen-tracker

> English version: [README.md](README.md)

## 概要

花粉飛散量と症状を記録・可視化する iOS アプリ。
Google Pollen API からリアルタイムの花粉データを取得し、毎日の症状強度と紐づけて管理する。

## 技術スタック

| レイヤー         | 技術                                          |
| ---------------- | --------------------------------------------- |
| モバイル         | React Native (Expo)                           |
| バックエンド     | Go / Echo v5 / AWS Lambda                     |
| データベース     | Turso (libSQL)                                |
| インフラ         | AWS CDK (TypeScript) / API Gateway (HTTP API) |
| 認証             | API Key (`x-api-key`)                         |
| CI/CD            | GitHub Actions / EAS Build                    |
| 花粉データ       | Google Pollen API                             |
| ジオコーディング | Google Places API                             |

## ディレクトリ構成

```text
pollen-tracker/
├── backend/                  # Go Lambda (Echo v5)
│   ├── cmd/lambda/           # Lambda エントリポイント
│   ├── internal/
│   │   └── gen/              # oapi-codegen 生成コード (手動編集禁止)
│   ├── oas/                  # OpenAPI 仕様 (リソースごとに分割)
│   │   ├── components/       # スキーマ・セキュリティ定義
│   │   ├── config/           # oapi-codegen 設定ファイル
│   │   ├── paths/            # リソースごとのパス定義
│   │   ├── oapi-base.yaml    # $ref リンク付きルート仕様
│   │   └── oapi.yaml         # バンドル済み仕様 (自動生成)
│   ├── tools/
│   │   └── migrations/       # DB マイグレーション (goose)
│   │       ├── cmd/          # マイグレーション実行ラッパー (go run)
│   │       └── sql/          # goose SQL ファイル
│   └── Makefile
│
├── mobile/                   # React Native (Expo)
│   ├── assets/               # アイコン・スプラッシュ画像
│   └── src/                  # Feature-based ソースコード
│
├── infra/                    # AWS CDK (TypeScript)
│
├── .github/
│   ├── workflows/            # CI/CD ワークフロー
│   └── PULL_REQUEST_TEMPLATE.md
│
└── docs/
    ├── README.md
    └── README.ja.md
```

## セットアップ手順

### 前提条件

- Go 1.26.3 ([asdf](https://asdf-vm.com/) で管理)
- Node.js 24.13.0 (asdf で管理)
- AWS CDK CLI (`npm install -g aws-cdk`)
- Expo CLI (`npx expo`)
- Docker (ローカル開発用)

### Backend

開発ツールのインストール

```bash
cd backend
make setup-tools
```

API サーバ起動

```bash
make up
```

DB マイグレーション

```bash
# マイグレーション実行
make migrate-up

# マイグレーションファイルの新規作成
make migrate-create name=<migration_name>
```

コード生成 (OpenAPI 仕様 → Go)

```bash
make gen-api
```

### Frontend

```bash
cd mobile
npx expo start
```

### Infrastructure

```bash
cd infra
cdk bootstrap   # 初回のみ
cdk diff        # 変更内容の確認
cdk deploy      # インフラ適用
```
