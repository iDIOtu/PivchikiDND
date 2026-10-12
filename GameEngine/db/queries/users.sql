-- name: CreateUser :one
INSERT INTO users DEFAULT VALUES
RETURNING id, created_at;

-- name: GetUser :one
SELECT id, created_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, created_at
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1;
