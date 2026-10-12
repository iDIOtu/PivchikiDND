
-- name: CreateEvent :one
INSERT INTO events (
    game_id,
    type,
    actor_id,
    target_id,
    payload
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, game_id, type, actor_id, target_id,
          payload, created_at;

-- name: GetEvent :one
SELECT id, game_id, type, actor_id, target_id,
       payload, created_at
FROM events
WHERE id = $1;

-- name: ListEventsByGame :many
SELECT id, game_id, type, actor_id, target_id,
       payload, created_at
FROM events
WHERE game_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: DeleteEvent :execrows
DELETE FROM events
WHERE id = $1;
