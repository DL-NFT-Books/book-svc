-- +migrate Up

alter table book
add column chain_id varchar not null default 0;

-- +migrate Down

alter table book
delete column chain_id;

