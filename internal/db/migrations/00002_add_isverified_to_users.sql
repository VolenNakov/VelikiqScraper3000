-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN is_validated BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN is_validated;
-- +goose StatementEnd
