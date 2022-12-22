-- +migrate Up

alter table book
add column network varchar not null default '';

-- +migrate Down

drop table key_value;

