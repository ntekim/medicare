-- name: CreateUser :one
INSERT INTO users (first_name, last_name, password_hash, email, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: ListDoctors :many
SELECT * FROM users WHERE role = 'doctor';