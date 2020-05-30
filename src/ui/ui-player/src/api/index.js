import Axios from "axios";

const axios = Axios.create({
    baseURL: process.env.REACT_APP_API_URL,
    timeout: 2000,
    headers: {
        'Accept': 'application/json',
    },
})

const distinctListGet = ({field, filter}) => axios.get(`/media/${field}/distinct`, {
    params: filter,
});

export {
    distinctListGet,
}