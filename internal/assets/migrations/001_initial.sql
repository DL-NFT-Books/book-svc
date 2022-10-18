-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(64) NOT NULL,
    description VARCHAR(500) NOT NULL,
    price VARCHAR(30) NOT NULL,
    banner JSONB NOT NULL,
    file JSONB NOT NULL
);
-- +migrate Down
DROP TABLE book;
