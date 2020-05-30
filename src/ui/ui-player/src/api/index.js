import Axios from "axios";

const axios = Axios.create({
    baseURL: 'http://localhost:8000',
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