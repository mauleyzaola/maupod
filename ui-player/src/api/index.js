import Axios from "axios";
import querystring from 'querystring';

const axios = Axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    timeout: 2000,
    headers: {
        'Accept': 'application/json',
    },
})

const audioScan = data => axios.post(`/audio/scan`, data);

const albumViewList = data => axios.get(`/media/albums`, {
    params: data,
})

const decodeURL = search => querystring.decode(search.replace('?',''));

const distinctListGet = ({field, filter}) => axios.get(`/media/${field}/distinct`, {
    params: filter,
});
const ipcCommand = data => axios.post(`/ipc`, data);
const genreList = data => axios.get(`/genres`, { params: data});
const genreArtworkList = data => axios.get(`/genres/artwork`, { params: data});
const mediaList = (data) => axios.get(`/media`, {
    params: data,
});
const queueAdd = ({media, index = -1, named_position}) => axios.post(`/queue`, {media, index, named_position});
const objectToQueryString = data => querystring.stringify(data);

export {
    audioScan,
    albumViewList,
    decodeURL,
    distinctListGet,
    ipcCommand,
    genreArtworkList,
    genreList,
    mediaList,
    objectToQueryString,
    queueAdd,
}