
-- +migrate Up

create table media_queue(
	id text primary key ,
	position int not null,
	media_id text not null
);

-- +migrate Down

drop table media_queue;