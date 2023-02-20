-- +migrate Up

ALTER TABLE book
ADD column IF NOT EXISTS floor_price VARCHAR(30) DEFAULT '0';

-- +migrate Down

ALTER TABLE book
DROP COLUMN IF EXISTS floor_price;