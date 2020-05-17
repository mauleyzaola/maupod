
-- +migrate Up

create table media (
    id text primary key,
    sha text not null,
    location text not null unique,
    file_extension text not null ,
    format text not null ,
    file_size bigint not null ,
    duration double precision not null,
    overall_bit_rate_mode text not null,
    overall_bit_rate bigint not null,
    stream_size bigint not null,
    album text not null,
    track text not null,
    title text not null,
    track_position bigint not null,
    performer text not null,
    genre text not null,
    recorded_date bigint not null,
    file_modified_date timestamp not null,
    comment text not null,
    channels text not null,
    channel_positions text not null,
    channel_layout text not null,
    sampling_rate bigint not null,
    sampling_count bigint not null,
    bit_depth bigint not null,
    compression_mode text not null,
    encoded_library text not null,
    encoded_library_name text not null,
    encoded_library_version text not null,
    bit_rate_mode text not null,
    bit_rate bigint not null,
    last_scan timestamp without time zone not null,
    modified_date timestamp with time zone not null
);

-- +migrate Down

drop table media;