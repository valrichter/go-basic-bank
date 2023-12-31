// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one

INSERT INTO
    users (
        username,
        hashed_password,
        full_name,
        email
    )
VALUES ($1, $2, $3, $4)
RETURNING username, hashed_password, full_name, email, password_chaged_at, created_at, is_email_verified
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChagedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const getUser = `-- name: GetUser :one

SELECT username, hashed_password, full_name, email, password_chaged_at, created_at, is_email_verified FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChagedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one

UPDATE users
SET
    hashed_password = coalesce(
        $1,
        hashed_password
    ),
    password_chaged_at = coalesce(
        $2,
        password_chaged_at
    ),
    full_name = coalesce(
        $3,
        full_name
    ),
    email = coalesce($4, email)
WHERE
    username = $5
RETURNING username, hashed_password, full_name, email, password_chaged_at, created_at, is_email_verified
`

type UpdateUserParams struct {
	HashedPassword   sql.NullString `json:"hashed_password"`
	PasswordChagedAt sql.NullTime   `json:"password_chaged_at"`
	FullName         sql.NullString `json:"full_name"`
	Email            sql.NullString `json:"email"`
	Username         string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChagedAt,
		arg.FullName,
		arg.Email,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChagedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}
