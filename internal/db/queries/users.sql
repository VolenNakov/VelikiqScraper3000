-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: VerifyUser :one
UPDATE users SET is_verified = ? WHERE id = ? RETURNING is_verified;

-- name: CreateUser :execlastid
INSERT INTO users (username, password_hash) VALUES (?, ?);

-- name: GetUnverifiedUsers :many
SELECT * FROM users WHERE is_verified = 0 ORDER BY id