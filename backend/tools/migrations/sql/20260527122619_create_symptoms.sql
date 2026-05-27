-- +goose Up
-- 症状ログテーブル: 日別のアレルギー症状強度 (1=軽微, 5=重症) を記録する
CREATE TABLE IF NOT EXISTS `symptoms` (
    `id`         INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, -- 症状ログ ID
    `date`       TEXT    NOT NULL UNIQUE,                    -- 記録日 (ISO 8601: YYYY-MM-DD)
    `sneezing`   INTEGER NOT NULL CHECK (`sneezing`   BETWEEN 1 AND 5), -- くしゃみ強度
    `runny_nose` INTEGER NOT NULL CHECK (`runny_nose` BETWEEN 1 AND 5), -- 鼻水強度
    `eye_itch`   INTEGER NOT NULL CHECK (`eye_itch`   BETWEEN 1 AND 5), -- 目のかゆみ強度
    `created_at` TEXT    NOT NULL DEFAULT (STRFTIME('%Y-%m-%dT%H:%M:%fZ', 'NOW')), -- 作成日時
    `updated_at` TEXT    NOT NULL DEFAULT (STRFTIME('%Y-%m-%dT%H:%M:%fZ', 'NOW'))  -- 更新日時
);

-- +goose Down
DROP TABLE IF EXISTS `symptoms`;
