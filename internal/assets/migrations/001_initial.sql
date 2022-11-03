-- +migrate Up
CREATE TABLE book(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(64) NOT NULL,
    symbol varchar(8) not null,
    description VARCHAR(500) NOT NULL,
    price VARCHAR(30) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    contract_name VARCHAR(30) NOT NULL,
    contract_version VARCHAR(20) NOT NULL,
    banner JSONB NOT NULL,
    file JSONB NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT 'f',
    last_block bigint
);
-- +migrate Down
DROP TABLE book;
