-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN email TO username;
DROP INDEX IF EXISTS idx_users_email;
CREATE INDEX IF NOT EXISTS idx_users_username on users(username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN username TO email;
DROP INDEX IF EXISTS idx_users_username;
CREATE INDEX IF NOT EXISTS idx_users_email on users(email);
-- +goose StatementEnd
