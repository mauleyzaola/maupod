import { combineReducers } from "redux";

const queueRows = (state = [], action) => {
    switch (action.type){
        case 'QUEUE_ADD':
            return [
                ...state,
                action.media,
            ];
        default:
            return state;
    }
}

export default combineReducers({
    queueRows,
});