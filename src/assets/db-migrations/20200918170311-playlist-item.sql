
 -- +migrate Up

 create table playlist_item(
 		id          text primary key, 
 		playlist_id text not null REFERENCES playlist(id), 
 		position    int not null, 
 		media_id    text not null REFERENCES media(id)
 		);

-- +migrate Down

drop table playlist_item;