-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN is_validated TO is_verified;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN is_verified TO is_validated;
-- +goose StatementEnd
