
-- +migrate Up

create or replace view view_genres as

select genre,
       cast(count(distinct performer) as double precision) as performer_count,
       cast(count(distinct album_identifier) as double precision) as album_count,
       cast(sum(duration) as double precision) as duration,
       cast(count(*) as double precision) as total
from media
where length(album_identifier) > 0
group by genre;

-- +migrate Down

drop view if exists view_genres;