import API from '../api';

function addQueue (queue) {
    return {
        type: ADD_QUEUE,
        queue,
    }
}

function removeQueue (index) {
    return {
        type: REMOVE_QUEUE,
        index,
    }
}

function loadQueue(rows) {
    return {
        type: LOAD_QUEUE,
        rows,
    }
}

export function handleDeleteQueue (queue) {
    return (dispatch) => {
        dispatch(removeQueue(queue.index));
        return API.queueRemove(queue.index)
            .catch(() => {
                dispatch(addQueue(queue));
                console.warn('An error occurred. Try again.');
            })
    }
}

export function handleAddQueue (queue) {
    return (dispatch) => {
        return API.queueAdd(queue)
            .then((media) => {
                dispatch(addQueue(queue));
            })
            .catch(() => console.warn('There was an error. Try again.'));
    }
}

export function handleLoadQueue() {
    return(dispatch) => {
        return API.queueList()
            .then(response => dispatch(loadQueue(response.data.rows || [])))
            .catch(() => console.warn('There was an error. Try again.'));
    }
}

export const ADD_QUEUE = 'ADD_QUEUE';
export const REMOVE_QUEUE = 'REMOVE_QUEUE';
export const LOAD_QUEUE = 'LOAD_QUEUE';
