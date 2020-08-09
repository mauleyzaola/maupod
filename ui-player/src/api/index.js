import Axios from "axios";
import querystring from 'querystring';

const axios = Axios.create({
    baseURL: process.env.REACT_APP_MAUPOD_API,
    timeout: 2000,
    headers: {
        'Accept': 'application/json',
    },
})

const cleanMedia = media => {
    const result = Object.assign({}, media);
    result.recorded_date = result.recorded_date || 0;
    return result;
}


const audioScan = data => axios.post(`/audio/scan`, data);

const albumViewList = data => axios.get(`/media/albums`, {
    params: data,
})

const decodeURL = search => querystring.decode(search.replace('?',''));
const directoryRead = async data => {
    const response = await axios.post(`/file-browser/directory`, data);
    return response.data.files || [];
}
const distinctListGet = ({field, filter}) => axios.get(`/media/${field}/distinct`, {
    params: filter,
});
const ipcCommand = data => {
    data.media = cleanMedia(data.media);
    return axios.post(`/ipc`, data);
}
const genreList = data => axios.get(`/genres`, { params: data});
const genreArtworkList = data => axios.get(`/genres/artwork`, { params: data});
const mediaList = (data) => axios.get(`/media`, {
    params: data,
});
const queueAdd = ({media, index = -1, named_position}) => axios.post(`/queue`, {media: cleanMedia(media), index, named_position});
const objectToQueryString = data => querystring.stringify(data);
const spectrumImage = (id) => axios.get(`/media/${id}/spectrum`);

export {
    audioScan,
    albumViewList,
    decodeURL,
    directoryRead,
    distinctListGet,
    ipcCommand,
    genreArtworkList,
    genreList,
    mediaList,
    objectToQueryString,
    queueAdd,
    spectrumImage,
}