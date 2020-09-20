
-- +migrate Up

create table playlist (
    id text primary key,
    name text not null unique
);

create table playlist_item (
    id text primary key,
    playlist_id text not null references playlist(id),
    position int not null,
    media_id text not null references media(id)
);

-- +migrate Down

drop table playlist_item;
drop table playlist;