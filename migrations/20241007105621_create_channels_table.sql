-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS yt_channels (
    id BIGSERIAL PRIMARY KEY,
    channel_id VARCHAR NOT NULL,
    topic VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    subscriptions INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_id unique (channel_id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
