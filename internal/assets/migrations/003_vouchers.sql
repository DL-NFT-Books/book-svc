-- +migrate Up

ALTER TABLE book
ADD column voucher_token VARCHAR(60) DEFAULT '',
ADD column voucher_token_amount VARCHAR(40) DEFAULT '';

-- +migrate Down

ALTER TABLE book
DROP COLUMN voucher_token,
DROP COLUMN voucher_token_amount;