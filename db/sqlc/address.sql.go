// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: address.sql

package db

import (
	"context"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (
  city,
  line1,
  line2,
  postcode
) VALUES (
  $1, $2, $3, $4
) RETURNING id, city, line1, line2, postcode
`

type CreateAddressParams struct {
	City     string `json:"city"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Postcode string `json:"postcode"`
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, createAddress,
		arg.City,
		arg.Line1,
		arg.Line2,
		arg.Postcode,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.Line1,
		&i.Line2,
		&i.Postcode,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1
`

func (q *Queries) DeleteAddress(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAddress, id)
	return err
}

const getAddressForUpdate = `-- name: GetAddressForUpdate :one
SELECT id, city, line1, line2, postcode FROM address
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetAddressForUpdate(ctx context.Context, id int64) (Address, error) {
	row := q.db.QueryRowContext(ctx, getAddressForUpdate, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.Line1,
		&i.Line2,
		&i.Postcode,
	)
	return i, err
}

const getUserAddress = `-- name: GetUserAddress :one
SELECT id, city, line1, line2, postcode FROM address
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserAddress(ctx context.Context, id int64) (Address, error) {
	row := q.db.QueryRowContext(ctx, getUserAddress, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.Line1,
		&i.Line2,
		&i.Postcode,
	)
	return i, err
}

const listAddresses = `-- name: ListAddresses :many
SELECT id, city, line1, line2, postcode FROM address
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAddressesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAddresses(ctx context.Context, arg ListAddressesParams) ([]Address, error) {
	rows, err := q.db.QueryContext(ctx, listAddresses, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Address{}
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.City,
			&i.Line1,
			&i.Line2,
			&i.Postcode,
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

const updateAddress = `-- name: UpdateAddress :one
UPDATE address
SET city = $2, line1 = $3, line2 = $4, postcode = $5
WHERE id = $1
RETURNING id, city, line1, line2, postcode
`

type UpdateAddressParams struct {
	ID       int64  `json:"id"`
	City     string `json:"city"`
	Line1    string `json:"line1"`
	Line2    string `json:"line2"`
	Postcode string `json:"postcode"`
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRowContext(ctx, updateAddress,
		arg.ID,
		arg.City,
		arg.Line1,
		arg.Line2,
		arg.Postcode,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.City,
		&i.Line1,
		&i.Line2,
		&i.Postcode,
	)
	return i, err
}
