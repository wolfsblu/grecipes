-- name: GetPasswordResetTokenByUser :one
SELECT *
FROM password_resets
WHERE user_id = ?
LIMIT 1;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES (?, ?)
RETURNING *;

-- name: CreatePasswordResetToken :one
INSERT INTO password_resets (user_id, token)
VALUES (?, ?)
RETURNING *;