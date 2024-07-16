-- name: CreateUser :one
INSERT INTO users (email, password, name,verified)
VALUES ($1, $2, $3,$4)
RETURNING id, email, name, created_at, updated_at, verified;

-- name: GetUserByEmail :one
SELECT id, email, name, created_at, updated_at,verified
FROM users
WHERE email = $1;


