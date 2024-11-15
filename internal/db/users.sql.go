// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (email, password_hash) VALUES (?, ?)
`

type CreateUserParams struct {
	Email        string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Email, arg.PasswordHash)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, password_hash, created_at FROM users WHERE email = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}
