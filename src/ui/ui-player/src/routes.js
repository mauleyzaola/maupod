import {objectToQueryString} from "./api";

const linkAlbumList = media => `/media?${objectToQueryString({album:media.album, performer: media.performer} )}`;
const linkGenreList = media => `/media?${objectToQueryString({genre:media.genre} )}`;
const linkPerformerList = media => `/media?${objectToQueryString({performer:media.performer} )}`;

export {
    linkAlbumList,
    linkGenreList,
    linkPerformerList,
}