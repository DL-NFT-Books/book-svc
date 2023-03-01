-- +migrate Up

ALTER TABLE book
ADD column voucher_token VARCHAR(60) DEFAULT '0x0000000000000000000000000000000000000000',
ADD column voucher_token_amount VARCHAR(40) DEFAULT '0';

-- +migrate Down

ALTER TABLE book
DROP COLUMN voucher_token,
DROP COLUMN voucher_token_amount;