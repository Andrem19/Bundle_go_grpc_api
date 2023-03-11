// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const changePassword = `-- name: ChangePassword :one
UPDATE users
SET hashed_password = $2
WHERE email = $1
RETURNING id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at
`

type ChangePasswordParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) ChangePassword(ctx context.Context, arg ChangePasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, changePassword, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  first_name,
  second_name,
  email,
  hashed_password,
  roles
  )VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at
`

type CreateUserParams struct {
	Username       string   `json:"username"`
	FirstName      string   `json:"first_name"`
	SecondName     string   `json:"second_name"`
	Email          string   `json:"email"`
	HashedPassword string   `json:"hashed_password"`
	Roles          []string `json:"roles"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.FirstName,
		arg.SecondName,
		arg.Email,
		arg.HashedPassword,
		pq.Array(arg.Roles),
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1
`

func (q *Queries) DeleteUser(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, email)
	return err
}

const emailConfirmd = `-- name: EmailConfirmd :one
UPDATE users
SET email_confirm = $2
WHERE email = $1
RETURNING id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at
`

type EmailConfirmdParams struct {
	Email        string `json:"email"`
	EmailConfirm bool   `json:"email_confirm"`
}

func (q *Queries) EmailConfirmd(ctx context.Context, arg EmailConfirmdParams) (User, error) {
	row := q.db.QueryRowContext(ctx, emailConfirmd, arg.Email, arg.EmailConfirm)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at FROM users
WHERE email = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FirstName,
			&i.SecondName,
			&i.Email,
			&i.EmailConfirm,
			pq.Array(&i.Roles),
			&i.HashedPassword,
			&i.Account,
			&i.Address,
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

const updateUserEmail = `-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
RETURNING id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at
`

type UpdateUserEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserEmail, arg.ID, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const updateUserName = `-- name: UpdateUserName :one
UPDATE users
SET first_name = $2, second_name = $3
WHERE email = $1
RETURNING id, username, first_name, second_name, email, email_confirm, roles, hashed_password, account, address, created_at
`

type UpdateUserNameParams struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserName, arg.Email, arg.FirstName, arg.SecondName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.SecondName,
		&i.Email,
		&i.EmailConfirm,
		pq.Array(&i.Roles),
		&i.HashedPassword,
		&i.Account,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}