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
     contract_address VARCHAR(42) NOT NULL,
     chain_id bigint not null default 0
);
ALTER TABLE book_network ADD UNIQUE ("book_id", "chain_id");

-- +migrate Down
DROP TABLE book;
DROP TABLE book_network;
