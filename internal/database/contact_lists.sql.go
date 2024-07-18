// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: contact_lists.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addContactToList = `-- name: AddContactToList :exec
INSERT INTO contact_lists (contact_id, list_id)
VALUES ($1, $2)
`

type AddContactToListParams struct {
	ContactID uuid.UUID
	ListID    uuid.UUID
}

func (q *Queries) AddContactToList(ctx context.Context, arg AddContactToListParams) error {
	_, err := q.db.ExecContext(ctx, addContactToList, arg.ContactID, arg.ListID)
	return err
}

const getContactsByListID = `-- name: GetContactsByListID :many
SELECT c.id, c.first_name, c.last_name, c.subscribed, c.blocklisted, c.email, c.whatsapp, c.landline_number, c.last_changed, c.date_added
FROM contacts c
JOIN contact_lists cl ON c.id = cl.contact_id
WHERE cl.list_id = $1
`

func (q *Queries) GetContactsByListID(ctx context.Context, listID uuid.UUID) ([]Contact, error) {
	rows, err := q.db.QueryContext(ctx, getContactsByListID, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contact
	for rows.Next() {
		var i Contact
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Subscribed,
			&i.Blocklisted,
			&i.Email,
			&i.Whatsapp,
			&i.LandlineNumber,
			&i.LastChanged,
			&i.DateAdded,
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

const removeContactFromList = `-- name: RemoveContactFromList :exec
DELETE FROM contact_lists
WHERE contact_id = $1 AND list_id = $2
`

type RemoveContactFromListParams struct {
	ContactID uuid.UUID
	ListID    uuid.UUID
}

func (q *Queries) RemoveContactFromList(ctx context.Context, arg RemoveContactFromListParams) error {
	_, err := q.db.ExecContext(ctx, removeContactFromList, arg.ContactID, arg.ListID)
	return err
}
