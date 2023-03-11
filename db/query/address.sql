-- name: CreateAddress :one
INSERT INTO address (
  city,
  line1,
  line2,
  postcode
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserAddress :one
SELECT * FROM address
WHERE id = $1 LIMIT 1;

-- name: GetAddressForUpdate :one
SELECT * FROM address
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAddresses :many
SELECT * FROM address
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAddress :one
UPDATE address
SET city = $2, line1 = $3, line2 = $4, postcode = $5
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1;