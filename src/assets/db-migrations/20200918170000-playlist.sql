 
 -- +migrate Up
 
 create table playlist(
 		id text primary key,
 		name text not null
 		);

-- +migrate Down

drop table playlist;
