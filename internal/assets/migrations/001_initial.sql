-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR(500) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    banner JSONB NOT NULL,
    file JSONB NOT NULL
);

CREATE TABLE book_network(
     book_id BIGSERIAL REFERENCES book(id) ON DELETE CASCADE,
     token_id bigint,
     chain_id bigint not null default 0
);
-- +migrate Down
DROP TABLE book;
DROP TABLE book_network;
