-- +goose Up
-- 症状ログテーブル: 日別のアレルギー症状強度 (1=軽微, 5=重症) を記録する
CREATE TABLE IF NOT EXISTS symptoms (
    id              TEXT    NOT NULL PRIMARY KEY,                                     -- UUID v7
    date            TEXT    NOT NULL UNIQUE,                                          -- ISO 8601: YYYY-MM-DD
    sneezing        INTEGER NOT NULL CHECK (sneezing     BETWEEN 1 AND 5),           -- くしゃみ強度
    runny           INTEGER NOT NULL CHECK (runny        BETWEEN 1 AND 5),           -- 鼻水強度
    itchy           INTEGER NOT NULL CHECK (itchy        BETWEEN 1 AND 5),           -- 目のかゆみ強度
    pollen_level    INTEGER NOT NULL CHECK (pollen_level BETWEEN 1 AND 5),           -- 花粉レベル
    took_medication INTEGER NOT NULL DEFAULT 0,                                      -- 服薬: 0=false, 1=true
    note            TEXT    NOT NULL DEFAULT '',                                     -- フリーメモ (最大 200 文字)
    created_at      TEXT    NOT NULL DEFAULT (STRFTIME('%Y-%m-%dT%H:%M:%fZ', 'NOW')), -- 作成日時
    updated_at      TEXT    NOT NULL DEFAULT (STRFTIME('%Y-%m-%dT%H:%M:%fZ', 'NOW'))  -- 更新日時
);

-- +goose Down
DROP TABLE IF EXISTS symptoms;
