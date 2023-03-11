-- name: CreateUser :one
INSERT INTO users (
  username,
  first_name,
  second_name,
  email,
  hashed_password,
  roles,
  account,
  address
  )VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE email = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserName :one
UPDATE users
SET first_name = $2, second_name = $3
WHERE email = $1
RETURNING *;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE id = $1
RETURNING *;

-- name: EmailConfirmd :one
UPDATE users
SET email_confirm = $2
WHERE email = $1
RETURNING *;

-- name: ChangePassword :one
UPDATE users
SET hashed_password = $2
WHERE email = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;