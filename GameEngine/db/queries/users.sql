-- name: CreateUser :one
INSERT INTO users DEFAULT VALUES
RETURNING id, created_at;

-- name: GetUser :one
SELECT id, created_at
FROM users
WHERE id = $1;

-- name: CreateGame :one
INSERT INTO games (
    user_id,
    name,
    setting_description,
    ruleset
)
VALUES ($1, $2, $3, $4)
RETURNING
    id,
    user_id,
    name,
    setting_description,
    ruleset,
    created_at,
    updated_at;

-- name: GetGame :one
SELECT
    id,
    user_id,
    name,
    setting_description,
    ruleset,
    created_at,
    updated_at
FROM games
WHERE id = $1;

-- name: CreateEntity :one
INSERT INTO entities (
    game_id,
    type,
    name,
    data
)
VALUES ($1, $2, $3, $4)
RETURNING
    id,
    game_id,
    type,
    name,
    data,
    created_at,
    updated_at;

-- name: CreateEvent :one
INSERT INTO events (
    game_id,
    type,
    actor_id,
    target_id,
    payload
)
VALUES ($1, $2, $3, $4, $5)
RETURNING
    id,
    game_id,
    type,
    actor_id,
    target_id,
    payload,
    created_at;

-- name: ListGameEvents :many
SELECT
    id,
    game_id,
    type,
    actor_id,
    target_id,
    payload,
    created_at
FROM events
WHERE game_id = $1
ORDER BY created_at DESC, id DESC
LIMIT $2;