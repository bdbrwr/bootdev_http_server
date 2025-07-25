// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING id, created_at, updated_at, email, password, is_chirpy_red
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.IsChirpyRed,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, password, is_chirpy_red FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.IsChirpyRed,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET email = $2, password = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, created_at, updated_at, email, password, is_chirpy_red
`

type UpdateUserParams struct {
	ID       uuid.UUID
	Email    string
	Password string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.IsChirpyRed,
	)
	return i, err
}

const upgradeUserToRed = `-- name: UpgradeUserToRed :one
UPDATE users SET is_chirpy_red = TRUE
WHERE id = $1
RETURNING id, created_at, updated_at, email, password, is_chirpy_red
`

func (q *Queries) UpgradeUserToRed(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, upgradeUserToRed, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.Password,
		&i.IsChirpyRed,
	)
	return i, err
}
