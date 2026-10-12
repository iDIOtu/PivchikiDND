
-- name: CreateEntity :one
INSERT INTO entities (
    game_id,
    type,
    name,
    data
)
VALUES ($1, $2, $3, $4)
RETURNING id, game_id, type, name, data,
          created_at, updated_at;

-- name: GetEntity :one
SELECT id, game_id, type, name, data,
       created_at, updated_at
FROM entities
WHERE id = $1;

-- name: ListEntitiesByGame :many
SELECT id, game_id, type, name, data,
       created_at, updated_at
FROM entities
WHERE game_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdateEntity :one
UPDATE entities
SET type = $2,
    name = $3,
    data = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, game_id, type, name, data,
          created_at, updated_at;

-- name: DeleteEntity :execrows
DELETE FROM entities
WHERE id = $1;
