
-- +migrate Up

create table media_event (
    id text primary key ,
    sha text not null,
    media_id text not null,
    ts timestamp with time zone not null,
    event int not null
);

-- +migrate Down

drop table media_event;