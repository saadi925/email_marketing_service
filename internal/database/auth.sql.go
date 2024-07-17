// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPasswordResetToken = `-- name: CreatePasswordResetToken :exec
INSERT INTO password_reset_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
`

type CreatePasswordResetTokenParams struct {
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}

func (q *Queries) CreatePasswordResetToken(ctx context.Context, arg CreatePasswordResetTokenParams) error {
	_, err := q.db.ExecContext(ctx, createPasswordResetToken, arg.UserID, arg.Token, arg.ExpiresAt)
	return err
}

const deletePasswordResetToken = `-- name: DeletePasswordResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = $1
`

func (q *Queries) DeletePasswordResetToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deletePasswordResetToken, token)
	return err
}

const getPasswordResetToken = `-- name: GetPasswordResetToken :one
SELECT id, user_id, token, expires_at
FROM password_reset_tokens
WHERE token = $1
`

type GetPasswordResetTokenRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}

func (q *Queries) GetPasswordResetToken(ctx context.Context, token string) (GetPasswordResetTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getPasswordResetToken, token)
	var i GetPasswordResetTokenRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.ExpiresAt,
	)
	return i, err
}
