-- +goose Up

CREATE TABLE stats (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    score int NOT NULL
);

-- +goose Down

DROP TABLE stats;