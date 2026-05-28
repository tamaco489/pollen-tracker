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
