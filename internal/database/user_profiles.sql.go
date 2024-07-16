// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_profiles.sql

package database

import (
	"context"
	"database/sql"
)

const createUserProfile = `-- name: CreateUserProfile :one
INSERT INTO user_profiles (user_id, email, first_name, last_name, phone_number, company_name, website, street_address, zip_code, city, country)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, user_id, email, first_name, last_name, created_at, updated_at
`

type CreateUserProfileParams struct {
	UserID        int32
	Email         string
	FirstName     string
	LastName      string
	PhoneNumber   sql.NullString
	CompanyName   string
	Website       sql.NullString
	StreetAddress string
	ZipCode       string
	City          string
	Country       string
}

type CreateUserProfileRow struct {
	ID        int32
	UserID    int32
	Email     string
	FirstName string
	LastName  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

func (q *Queries) CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (CreateUserProfileRow, error) {
	row := q.db.QueryRowContext(ctx, createUserProfile,
		arg.UserID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.CompanyName,
		arg.Website,
		arg.StreetAddress,
		arg.ZipCode,
		arg.City,
		arg.Country,
	)
	var i CreateUserProfileRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProfileByUserID = `-- name: GetUserProfileByUserID :one
SELECT id, user_id, email, first_name, last_name, phone_number, company_name, website, street_address, zip_code, city, country, created_at, updated_at
FROM user_profiles
WHERE user_id = $1
`

func (q *Queries) GetUserProfileByUserID(ctx context.Context, userID int32) (UserProfile, error) {
	row := q.db.QueryRowContext(ctx, getUserProfileByUserID, userID)
	var i UserProfile
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.CompanyName,
		&i.Website,
		&i.StreetAddress,
		&i.ZipCode,
		&i.City,
		&i.Country,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
