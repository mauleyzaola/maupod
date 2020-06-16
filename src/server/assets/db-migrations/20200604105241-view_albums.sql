
-- +migrate Up

create or replace view view_albums as

select max(id) as id, album_identifier, max(album) as album, sum(cast(duration as int)) as duration,
       cast(avg(bit_rate) as integer) as bit_rate, max(performer) as performer,
       max(genre) as genre, cast(max(recorded_date) as integer) as recorded_date,
       cast(avg(sampling_rate) as integer) as sampling_rate, cast(max(track_name_total) as integer) as track_name_total,
       max(sha_image) as sha_image, max(format) as format, cast(sum(file_size) as integer) as file_size
from media
where album_identifier <> ''
group by album_identifier;

-- +migrate Down

drop view if exists view_albums;