-- name: GetSymptoms :many
SELECT
    "id",
    "date",
    "sneezing",
    "runny",
    "itchy",
    "pollen_level",
    "took_medication",
    "note",
    "created_at",
    "updated_at"
FROM "symptoms"
WHERE "date" BETWEEN sqlc.arg('from') AND sqlc.arg('to')
ORDER BY "date" DESC;

-- name: UpdateSymptom :one
UPDATE "symptoms"
SET
    "sneezing"        = sqlc.arg('sneezing'),
    "runny"           = sqlc.arg('runny'),
    "itchy"           = sqlc.arg('itchy'),
    "pollen_level"    = sqlc.arg('pollen_level'),
    "took_medication" = sqlc.arg('took_medication'),
    "note"            = sqlc.arg('note'),
    "updated_at"      = sqlc.arg('updated_at')
WHERE "id" = sqlc.arg('id')
RETURNING
    "id",
    "date",
    "sneezing",
    "runny",
    "itchy",
    "pollen_level",
    "took_medication",
    "note",
    "created_at",
    "updated_at";

-- name: GetWeeklyStats :many
SELECT
    strftime('%Y-W%W', "date")              AS period_key,
    MIN("date")                             AS start_date,
    MAX("date")                             AS end_date,
    ROUND(AVG(CAST("sneezing"     AS REAL)), 1) AS avg_sneezing,
    ROUND(AVG(CAST("runny"        AS REAL)), 1) AS avg_runny,
    ROUND(AVG(CAST("itchy"        AS REAL)), 1) AS avg_itchy,
    ROUND(AVG(CAST("pollen_level" AS REAL)), 1) AS avg_pollen_level,
    COUNT(*)                                AS count
FROM "symptoms"
WHERE "date" BETWEEN sqlc.arg('from') AND sqlc.arg('to')
GROUP BY strftime('%Y-W%W', "date")
ORDER BY period_key ASC;

-- name: GetMonthlyStats :many
SELECT
    strftime('%Y-%m', "date")               AS period_key,
    MIN("date")                             AS start_date,
    MAX("date")                             AS end_date,
    ROUND(AVG(CAST("sneezing"     AS REAL)), 1) AS avg_sneezing,
    ROUND(AVG(CAST("runny"        AS REAL)), 1) AS avg_runny,
    ROUND(AVG(CAST("itchy"        AS REAL)), 1) AS avg_itchy,
    ROUND(AVG(CAST("pollen_level" AS REAL)), 1) AS avg_pollen_level,
    COUNT(*)                                AS count
FROM "symptoms"
WHERE "date" BETWEEN sqlc.arg('from') AND sqlc.arg('to')
GROUP BY strftime('%Y-%m', "date")
ORDER BY period_key ASC;

-- name: GetSymptomPollenLevels :many
SELECT "pollen_level"
FROM "symptoms"
WHERE "sneezing" + "runny" + "itchy" > 0
    AND "pollen_level" > 0
ORDER BY "pollen_level" ASC;

-- name: InsertSymptom :exec
INSERT INTO "symptoms" (
    "id",
    "date",
    "sneezing",
    "runny",
    "itchy",
    "pollen_level",
    "took_medication",
    "note",
    "created_at",
    "updated_at"
) VALUES (
    sqlc.arg('id'),
    sqlc.arg('date'),
    sqlc.arg('sneezing'),
    sqlc.arg('runny'),
    sqlc.arg('itchy'),
    sqlc.arg('pollen_level'),
    sqlc.arg('took_medication'),
    sqlc.arg('note'),
    sqlc.arg('created_at'),
    sqlc.arg('updated_at')
);
