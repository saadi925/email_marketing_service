// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: email_logs.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createEmailLog = `-- name: CreateEmailLog :one
INSERT INTO email_logs (email_id, status, message)
VALUES ($1, $2, $3)
RETURNING id, email_id, status, message, created_at
`

type CreateEmailLogParams struct {
	EmailID uuid.UUID
	Status  string
	Message sql.NullString
}

func (q *Queries) CreateEmailLog(ctx context.Context, arg CreateEmailLogParams) (EmailLog, error) {
	row := q.db.QueryRowContext(ctx, createEmailLog, arg.EmailID, arg.Status, arg.Message)
	var i EmailLog
	err := row.Scan(
		&i.ID,
		&i.EmailID,
		&i.Status,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}

const getEmailLogsByEmailID = `-- name: GetEmailLogsByEmailID :many
SELECT id, email_id, status, message, created_at
FROM email_logs
WHERE email_id = $1
`

func (q *Queries) GetEmailLogsByEmailID(ctx context.Context, emailID uuid.UUID) ([]EmailLog, error) {
	rows, err := q.db.QueryContext(ctx, getEmailLogsByEmailID, emailID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmailLog
	for rows.Next() {
		var i EmailLog
		if err := rows.Scan(
			&i.ID,
			&i.EmailID,
			&i.Status,
			&i.Message,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
