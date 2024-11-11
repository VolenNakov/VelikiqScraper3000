-- name: GetUser :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (email, password_hash) VALUES (?, ?);
