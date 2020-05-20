
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
    modified_date timestamp without time zone not null,
    track_name_total bigint not null,
    album_performer text not null,
    audio_count bigint not null,
    bit_depth_string text not null,
    commercial_name text not null,
    complete_name text not null,
    count_of_audio_streams bigint not null,
    encoded_library_date text not null,
    file_name text not null,
    folder_name text not null,
    format_info text not null,
    format_url text not null,
    internet_media_type text not null,
    kind_of_stream text not null,
    part bigint not null,
    part_total bigint not null,
    stream_identifier bigint not null,
    writing_library text not null
);

-- +migrate Down

drop table media;