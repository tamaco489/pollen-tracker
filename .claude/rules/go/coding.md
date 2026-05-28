# Go コーディングルール

## パッケージエイリアス

- エイリアスはスネークケースで統一する (例: `domain_pollen`, `google_pollen`)
- **同一ファイル内で名前衝突が発生する場合のみ**付与する — 衝突がなければエイリアスなしで import する

## import グループ順序

空白行で区切り、以下の順序に従う。

```go
import (
    // 1. 標準パッケージ
    "context"
    "errors"

    // 2. 内部パッケージ (github.com/tamaco489/pollen-tracker/...)
    "github.com/tamaco489/pollen-tracker/backend/internal/gen"
    "github.com/tamaco489/pollen-tracker/backend/pkg/errors/sentinel"

    // 3. エイリアスを付与した外部・内部パッケージ
    openapi_types "github.com/oapi-codegen/runtime/types"
)
```
