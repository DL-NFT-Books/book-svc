-- +migrate Up

ALTER TABLE IF EXISTS book
    ADD COLUMN IF NOT EXISTS chain_id BIGINT NOT NULL DEFAULT 5; -- TODO: SET POLYGON CHAIN_ID (137)

CREATE INDEX idx_book_chain_id ON book (chain_id);

-- +migrate Down

DROP INDEX IF EXISTS idx_book_chain_id;
ALTER TABLE IF EXISTS book
    DROP COLUMN IF EXISTS chain_id;
