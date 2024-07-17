// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: email_verify.sql

package database

import (
	"context"
	"time"
)

const createEmailVerify = `-- name: CreateEmailVerify :exec
INSERT INTO email_verify (
    email,
    code,
    retry,
    last_attempt
) VALUES (
    $1, $2, $3, $4
)
`

type CreateEmailVerifyParams struct {
	Email       string
	Code        string
	Retry       int32
	LastAttempt time.Time
}

func (q *Queries) CreateEmailVerify(ctx context.Context, arg CreateEmailVerifyParams) error {
	_, err := q.db.ExecContext(ctx, createEmailVerify,
		arg.Email,
		arg.Code,
		arg.Retry,
		arg.LastAttempt,
	)
	return err
}

const emailVerifyExists = `-- name: EmailVerifyExists :one
SELECT EXISTS (
    SELECT 1
    FROM email_verify
    WHERE email = $1
)
`

func (q *Queries) EmailVerifyExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, emailVerifyExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getEmailVerifyByEmail = `-- name: GetEmailVerifyByEmail :one
SELECT
    email,
    code,
    retry,
    last_attempt
FROM email_verify
WHERE email = $1
`

func (q *Queries) GetEmailVerifyByEmail(ctx context.Context, email string) (EmailVerify, error) {
	row := q.db.QueryRowContext(ctx, getEmailVerifyByEmail, email)
	var i EmailVerify
	err := row.Scan(
		&i.Email,
		&i.Code,
		&i.Retry,
		&i.LastAttempt,
	)
	return i, err
}

const updateEmailVerify = `-- name: UpdateEmailVerify :exec
UPDATE email_verify
SET
    code = $2,
    retry = $3,
    last_attempt = $4
WHERE email = $1
`

type UpdateEmailVerifyParams struct {
	Email       string
	Code        string
	Retry       int32
	LastAttempt time.Time
}

func (q *Queries) UpdateEmailVerify(ctx context.Context, arg UpdateEmailVerifyParams) error {
	_, err := q.db.ExecContext(ctx, updateEmailVerify,
		arg.Email,
		arg.Code,
		arg.Retry,
		arg.LastAttempt,
	)
	return err
}