-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email,
  nickname
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByusername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserBynickname :one
SELECT * FROM users
WHERE nickname = $1 LIMIT 1;

-- name: GetUserByemail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE username = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY create_at
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

-- name: UpdateNickname :one
UPDATE users
set nickname = $2
WHERE username = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
set hashed_password = $2
WHERE username = $1
RETURNING *;

-- name: UpdateEmail :one
UPDATE users
set email = $2
WHERE username = $1
RETURNING *;

