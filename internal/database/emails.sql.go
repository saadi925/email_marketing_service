// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: emails.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createEmail = `-- name: CreateEmail :one
INSERT INTO emails (campaign_id, recipient_email)
VALUES ($1, $2)
RETURNING id, campaign_id, recipient_email, status, subscription_id, sent_at, created_at, updated_at
`

type CreateEmailParams struct {
	CampaignID     uuid.NullUUID
	RecipientEmail string
}

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (Email, error) {
	row := q.db.QueryRowContext(ctx, createEmail, arg.CampaignID, arg.RecipientEmail)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.CampaignID,
		&i.RecipientEmail,
		&i.Status,
		&i.SubscriptionID,
		&i.SentAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEmailByID = `-- name: GetEmailByID :one
SELECT id, campaign_id, recipient_email, status, subscription_id, sent_at, created_at, updated_at
FROM emails
WHERE id = $1
`

func (q *Queries) GetEmailByID(ctx context.Context, id uuid.UUID) (Email, error) {
	row := q.db.QueryRowContext(ctx, getEmailByID, id)
	var i Email
	err := row.Scan(
		&i.ID,
		&i.CampaignID,
		&i.RecipientEmail,
		&i.Status,
		&i.SubscriptionID,
		&i.SentAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEmailsByCampaignID = `-- name: GetEmailsByCampaignID :many
SELECT id, campaign_id, recipient_email, status, subscription_id, sent_at, created_at, updated_at
FROM emails
WHERE campaign_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetEmailsByCampaignID(ctx context.Context, campaignID uuid.NullUUID) ([]Email, error) {
	rows, err := q.db.QueryContext(ctx, getEmailsByCampaignID, campaignID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Email
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.ID,
			&i.CampaignID,
			&i.RecipientEmail,
			&i.Status,
			&i.SubscriptionID,
			&i.SentAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateEmailStatus = `-- name: UpdateEmailStatus :exec
UPDATE emails
SET status = $2
WHERE id = $1
`

type UpdateEmailStatusParams struct {
	ID     uuid.UUID
	Status sql.NullString
}

func (q *Queries) UpdateEmailStatus(ctx context.Context, arg UpdateEmailStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateEmailStatus, arg.ID, arg.Status)
	return err
}
