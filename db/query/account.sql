-- name: CreateAccount :one
INSERT INTO account (
  balance
) VALUES (
  $1
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetUserAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: SubAccBalance :one
UPDATE account
SET balance = balance - $2
WHERE id = $1
RETURNING *;

-- name: AddAccBalance :one
UPDATE account
SET balance = balance + $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1;