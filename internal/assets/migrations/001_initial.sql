-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price INTEGER NOT NULL,
    banner JSONB NOT NULL,
    file JSONB NOT NULL
);
-- +migrate Down
DROP TABLE book;
