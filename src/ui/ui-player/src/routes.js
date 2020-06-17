import {objectToQueryString} from "./api";

const linkAlbumList = media => `/albums?${objectToQueryString({format:media.format, performer: media.performer, genre: media.genre,} )}`;
const linkGenreList = media => `/media?${objectToQueryString({genre:media.genre} )}`;
const linkPerformerList = media => `/media?${objectToQueryString({performer:media.performer} )}`;
const linkAlbumSongList = media => `/media?${objectToQueryString({album_identifier:media.album_identifier} )}`;

export {
    linkAlbumList,
    linkGenreList,
    linkPerformerList,
    linkAlbumSongList,
}