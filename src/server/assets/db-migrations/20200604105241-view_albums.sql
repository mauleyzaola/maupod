
-- +migrate Up

create or replace view view_albums as

select max(id) as id, album_identifier, max(album) as album, sum(duration) as duration, avg(bit_rate) as bit_rate, max(performer) as performer,
       max(genre) as genre, max(recorded_date) as recorded_date,
       avg(sampling_rate) as sampling_rate, max(track_name_total) as track_name_total,
       max(sha_image) as sha_image
from media
where album_identifier <> ''
group by album_identifier;

-- +migrate Down

drop view if exists view_albums;