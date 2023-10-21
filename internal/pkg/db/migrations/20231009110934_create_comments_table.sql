-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    post_id    BIGINT REFERENCES posts (id) ON DELETE CASCADE,
    content    TEXT                  NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE       DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd
