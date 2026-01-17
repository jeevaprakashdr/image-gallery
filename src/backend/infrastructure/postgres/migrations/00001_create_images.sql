-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied.

CREATE TABLE IF NOT EXISTS images (
    id uuid NOT NULL,
    title text,
    tags text,
    CONSTRAINT images__pk PRIMARY KEY (id)
);

-- +goose Down
-- SQL in section 'Down' is executed when this migration is rolled back.

DROP TABLE IF EXISTS images;