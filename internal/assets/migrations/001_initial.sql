-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR(500) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    banner JSONB DEFAULT NULL,
    file JSONB DEFAULT NULL
);

CREATE TABLE book_network(
     book_id BIGSERIAL REFERENCES book(id) ON DELETE CASCADE,
     contract_address VARCHAR(42) NOT NULL,
     token_id bigint,
     deploy_status int8,
     chain_id bigint not null default 0
);
-- +migrate Down
DROP TABLE book;
DROP TABLE book_network;
