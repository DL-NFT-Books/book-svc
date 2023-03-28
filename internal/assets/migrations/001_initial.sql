-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    description VARCHAR(500) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    contract_address VARCHAR(42) NOT NULL,
    banner JSONB NOT NULL,
    file JSONB NOT NULL,
    token_id bigint,
    deploy_status int8,
    chain_id varchar not null default 0
);
-- +migrate Down
DROP TABLE book;
