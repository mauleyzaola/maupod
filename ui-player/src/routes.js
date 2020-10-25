import API from "./api";

const linkAlbumList = media => `/albums?${API.objectToQueryString({format:media.format, performer: media.performer, genre: media.genre,} )}`;
const linkAlbumView = media => `/album/${media.album_identifier}`;
const linkGenreList = media => `/albums?${API.objectToQueryString({genre:media.genre} )}`;
const linkMediaList = media => `/media?${API.objectToQueryString(media)}`;
const linkPerformerList = media => `/albums?${API.objectToQueryString({performer:media.performer} )}`;

export {
    linkAlbumList,
    linkAlbumView,
    linkGenreList,
    linkMediaList,
    linkPerformerList,
}