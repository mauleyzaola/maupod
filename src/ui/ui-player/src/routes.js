import {objectToQueryString} from "./api";

const linkAlbumList = media => `/albums?${objectToQueryString({format:media.format, performer: media.performer, genre: media.genre,} )}`;
const linkAlbumView = media => `/album?${objectToQueryString({album_identifier:media.album_identifier} )}`;
const linkGenreList = media => `/albums?${objectToQueryString({genre:media.genre} )}`;
const linkMediaList = media => `/media?${objectToQueryString(media)}`;
const linkPerformerList = media => `/albums?${objectToQueryString({performer:media.performer} )}`;

export {
    linkAlbumList,
    linkAlbumView,
    linkGenreList,
    linkMediaList,
    linkPerformerList,
}