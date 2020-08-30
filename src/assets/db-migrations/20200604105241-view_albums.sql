
-- +migrate Up

create or replace view view_albums as

select max(id) as id, album_identifier, max(album) as album, sum(cast(duration as int)) as duration,
       cast(avg(bit_rate) as bigint) as bit_rate, max(performer) as performer,
       max(genre) as genre, cast(max(recorded_date) as bigint) as recorded_date,
       cast(avg(sampling_rate) as bigint) as sampling_rate, cast(count(*) as bigint) as track_name_total,
       max(image_location) as image_location, max(format) as format, cast(sum(file_size) as bigint) as file_size
from media
where album_identifier <> ''
group by album_identifier,image_location;

-- +migrate Down

drop view if exists view_albums;