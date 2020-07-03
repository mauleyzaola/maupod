
-- +migrate Up

alter table media add album_identifier text not null default ('');
alter table media add is_compilation boolean not null default(false);

-- +migrate Down
