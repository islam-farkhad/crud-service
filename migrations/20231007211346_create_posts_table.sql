-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts
(
    id         BIGSERIAL PRIMARY KEY NOT NULL,
    content    TEXT                  NOT NULL DEFAULT '',
    likes      INT                   NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE       DEFAULT NOW() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
