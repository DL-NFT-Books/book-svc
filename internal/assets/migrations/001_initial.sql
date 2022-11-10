-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(64) NOT NULL,
    description VARCHAR(500) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price VARCHAR(30) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    contract_name VARCHAR(64) NOT NULL,
    contract_symbol varchar(8) not null,
    contract_version VARCHAR(32) NOT NULL,
    banner JSONB NOT NULL,
    file JSONB NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT 'f',
    token_id bigint,
    deploy_status int8,
    last_block bigint
);
-- +migrate Down
DROP TABLE book;
