# CLAUDE.md

このファイルは Claude Code (claude.ai/code) がこのリポジトリで作業する際のガイダンスを提供します。

@.claude/rules/github/commit-types.md
@.claude/rules/github/commit-subject.md
@.claude/rules/github/pr-description.md

## プロジェクト概要

花粉飛散量と毎日のアレルギー症状を記録・可視化する iOS アプリ。
Google Pollen API からリアルタイムデータを取得し、症状との相関を把握する。

## アーキテクチャ概要

- **フロントエンド**: React Native (Expo) + TypeScript — Feature-based 構成
- **バックエンド**: Go + Echo v5 + AWS Lambda (Clean Architecture)
- **データベース**: Turso (libSQL)
- **インフラ**: AWS CDK (TypeScript) + API Gateway (HTTP API)
- **CI/CD**: GitHub Actions

## 制約事項

> [!IMPORTANT]
>
> - **即時性は求めない。時間をかけてでも根拠に基づく正確なアウトプットを行う**
> - **公式ドキュメントや関連資料の調査はメインコンテキストを汚さないよう、別途調査用エージェントに委譲する**

- **コード変更前に必ずファイルを Read ツールで読む**
- **変更は diff 形式で提示し、承認 (y) を得てから実行する**
- **git commit はユーザーの承認を得てから実行する**
- 応答は日本語・簡潔・直接的
- コメントは「なぜ」が自明でない場合のみ書く（「何をしているか」は書かない）

## 禁止事項

- `rm -rf` の使用禁止 — ファイル削除は `rm -f` を使う
- 明示的な指示なしの変更禁止
- Git フック・署名のスキップ禁止 (`--no-verify`, `--no-gpg-sign`)
- `main` ブランチへの直接 push 禁止
- 自動生成ファイルの手動編集禁止:
  - `pollen/backend/internal/gen/*.gen.go` (oapi-codegen)
  - `pollen/backend/internal/infrastructure/db/sqlc/` (sqlc)
- 機密情報のハードコーディング禁止 (API キー, トークン, 接続情報)
- 絵文字の使用禁止 (明示的に求められた場合を除く)
