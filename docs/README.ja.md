# pollen-tracker

> English version: [README.md](README.md)

## 概要

花粉飛散量と症状を記録・可視化する iOS アプリ。
Google Pollen API からリアルタイムの花粉データを取得し、毎日の症状強度と紐づけて管理する。

## 技術スタック

| レイヤー         | 技術                                            |
| ---------------- | ----------------------------------------------- |
| モバイル         | React Native (Expo)                             |
| バックエンド     | Go / Echo v5 / AWS Lambda                       |
| データベース     | Turso (libSQL)                                  |
| インフラ         | AWS CDK (TypeScript) / API Gateway (HTTP API)   |
| 認証             | API Key (`x-api-key`)                           |
| CI/CD            | GitHub Actions / EAS Build                      |
| 花粉データ       | Google Pollen API                               |
| ジオコーディング | Google Places API                               |

## ディレクトリ構成

```text
pollen-tracker/
├── pollen/
│   ├── backend/          # Go Lambda (Echo v5)
│   │   ├── cmd/lambda/   # Lambda エントリポイント
│   │   ├── internal/     # ハンドラー・ドメインロジック・DB
│   │   ├── migrations/   # goose マイグレーション SQL
│   │   ├── pkg/          # 共通パッケージ (config, errors)
│   │   └── openapi.yaml  # OpenAPI 3.0 仕様書
│   │
│   ├── mobile/           # React Native (Expo)
│   │   ├── assets/       # アイコン・スプラッシュ画像
│   │   └── src/          # Feature-based ソースコード
│   │
│   └── infra/            # AWS CDK (TypeScript)
│
├── .github/
│   └── workflows/        # CI/CD ワークフロー
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

### バックエンド (ローカル)

```bash
cd pollen/backend
docker compose up
```

### フロントエンド (ローカル)

```bash
cd pollen/mobile
npx expo start
```

### インフラ

```bash
cd pollen/infra
cdk bootstrap   # 初回のみ
cdk diff        # 変更内容の確認
cdk deploy      # インフラ適用
```

## API エンドポイント

| メソッド | パス               | 概要                             |
| -------- | ------------------ | -------------------------------- |
| GET      | `/pollen`          | 座標をもとに花粉レベルを取得する |
| POST     | `/symptoms`        | 当日の症状を記録する             |
| GET      | `/symptoms`        | 記録済みの症状一覧を取得する     |
| PUT      | `/symptoms/{date}` | 指定日の症状を更新する           |
