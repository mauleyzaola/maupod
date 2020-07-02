
-- +migrate Up

create table media_event (
    id text primary key ,
    media_id text not null references media(id),
    ts timestamp with time zone not null,
    event int not null
);

-- +migrate Down

drop table media_event;