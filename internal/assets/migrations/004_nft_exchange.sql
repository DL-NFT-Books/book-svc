-- +migrate Up

ALTER TABLE book
    ADD column floor_price VARCHAR(30) DEFAULT '0';

-- +migrate Down

ALTER TABLE book
DROP COLUMN floor_price;