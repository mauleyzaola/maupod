import Axios from "axios";
import querystring from 'querystring';

const axios = Axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    timeout: 2000,
    headers: {
        'Accept': 'application/json',
    },
})

const decodeURL = url => querystring.decode(url.search.replace('?',''));

const distinctListGet = ({field, filter}) => axios.get(`/media/${field}/distinct`, {
    params: filter,
});

const objectToQueryString = data => querystring.stringify(data);

const mediaList = (data) => axios.get(`/media`, {
    params: data,
});

export {
    decodeURL,
    distinctListGet,
    mediaList,
    objectToQueryString,
}