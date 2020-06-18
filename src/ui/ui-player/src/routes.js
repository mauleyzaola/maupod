import {objectToQueryString} from "./api";

const linkAlbumList = media => `/albums?${objectToQueryString({format:media.format, performer: media.performer, genre: media.genre,} )}`;
const linkGenreList = media => `/albums?${objectToQueryString({genre:media.genre} )}`;
const linkPerformerList = media => `/albums?${objectToQueryString({performer:media.performer} )}`;
const linkAlbumView = media => `/album?${objectToQueryString({album_identifier:media.album_identifier} )}`;

export {
    linkAlbumList,
    linkGenreList,
    linkPerformerList,
    linkAlbumView,
}