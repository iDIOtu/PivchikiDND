
-- name: CreateGame :one
INSERT INTO games (
    user_id,
    name,
    setting_description,
    ruleset
)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, name, setting_description,
          ruleset, created_at, updated_at;

-- name: GetGame :one
SELECT id, user_id, name, setting_description,
       ruleset, created_at, updated_at
FROM games
WHERE id = $1;

-- name: ListGamesByUser :many
SELECT id, user_id, name, setting_description,
       ruleset, created_at, updated_at
FROM games
WHERE user_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateGame :one
UPDATE games
SET name = $2,
    setting_description = $3,
    ruleset = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, name, setting_description,
          ruleset, created_at, updated_at;

-- name: DeleteGame :execrows
DELETE FROM games
WHERE id = $1;
