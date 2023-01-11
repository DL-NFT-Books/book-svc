-- +migrate Up

ALTER TABLE book
ADD column voucher_token VARCHAR(60) DEFAULT null,
ADD column voucher_token_amount bigint;

-- +migrate Down

ALTER TABLE book
DROP COLUMN voucher_token,
DROP COLUMN voucher_token_amount;