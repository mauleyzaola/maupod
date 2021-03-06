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

const genreArtworkList = params => axios.get(`/genres/artwork`, { params });

const mediaList = (data) => axios.get(`/media`, {
    params: data,
});

const providerMetadataCovers = ({params}) => axios.get(`/providers/metadata/cover`, {params});

const providerMetadataCoverPut = ({params, data}) => axios.put(`/providers/metadata/cover/${params.album_identifier}`, data);

const queueAdd = ({media, index = -1, named_position}) => axios.post(`/queue`, {media: cleanMedia(media), index, named_position});

const queueList = () => axios.get(`/queue`);

const queueRemove = index => axios.delete(`/queue/${index}`);

const objectToQueryString = data => querystring.stringify(data);

const spectrumImage = (id) => axios.get(`/media/${id}/spectrum`);

///PlayList

const playListAdd = async ({data}) => {
    try {
        const response = await axios.post(`/playlists`,data);
        return response.data;
    } catch (e){
        return e;
    }
}

const playLists = async params => {
    try {
        const response = await axios.get(`/playlists`,params);
        return response.data;
    } catch (e){
        return e;
    }
}

const playListsGet = async ({id}) => {
    try {
        const response = await axios.get(`/playlists/${id}`)
        return response.data;
    } catch (e){
        return e;
    }
}

const playListsPut = async ({id, data}) => {
    try {
        const response = await axios.put(`/playlists/${id}`, data)
        return response.data;
    } catch (e){
        return e;
    }
}

const playListsDelete = async ({id}) => {
    try {
        const response = await axios.delete(`/playlists/${id}`)
        return response.data;
    } catch (e){
        return e;
    }
}

const playListItemPost = async ({id,data}) => {
    try {
        const response = await axios.post(`/playlists/${id}/items`,data)
        return response.data;
    } catch (e){
        return e;
    }
}

const playListItemsGet = async ({id}) => {
        try {
            const response = await axios.get(`/playlists/${id}/items`)
            return response.data;
        } catch (e){
            return e;
        }
}

const playListItemPut = async ({id, position, data}) => {
    try {
        const response = await axios.put(`/playlists/${id}/items/${position}`, data)
        return response.data;
    } catch (e){
        return e;
    }
}

const PlayListItemDelete = async ({id, position}) => {
        try {
            const response = await axios.delete(`/playlists/${id}/items/${position}`)
            return response.data;
        } catch (e){
            return e;
        }
}

const volumeChange = async (data) => {
        try {
            const response = await axios.post(`/volume`, data)
            return response.data;
        } catch (e){
            return e;
        }
}

export default {
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
    playListAdd,
    playListItemPost,
    playListItemsGet,
    playListItemPut,
    PlayListItemDelete,    
    playLists,
    playListsGet,
    playListsPut,
    playListsDelete,
    providerMetadataCovers,
    providerMetadataCoverPut,
    queueAdd,
    queueList,
    queueRemove,
    spectrumImage,
    volumeChange,
}