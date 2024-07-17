-- name: CreateUser :one
INSERT INTO users (email, password, name,verified)
VALUES ($1, $2, $3,$4)
RETURNING id, email, name, created_at, updated_at, verified;

-- name: GetUserByEmail :one
SELECT id, email, name, created_at, updated_at,verified, password
FROM users
WHERE email = $1;
-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2
WHERE id = $1;

-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, created_at;


-- name: GetRefreshTokenByToken :one
SELECT id, user_id, token, expires_at, created_at
FROM refresh_tokens
WHERE token = $1;


-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = $1;


-- name: VerifyUserByEmail :exec
UPDATE users
SET verified = true
WHERE email = $1;

-- name: DeleteEmailVerifyByEmail :exec
DELETE FROM email_verify
WHERE email = $1;
