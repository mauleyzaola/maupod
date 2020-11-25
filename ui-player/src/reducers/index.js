import { combineReducers } from "redux";
import {ADD_QUEUE, LOAD_QUEUE, REMOVE_QUEUE} from "../actions/queue";

const queues = (state = [], action) => {
    switch (action.type){
        case ADD_QUEUE:
            return [
                ...state,
                action.queue,
            ];
        case REMOVE_QUEUE:
            const newState = [];
            state.forEach((val,index) => {
                if(index !== action.index){
                    newState.push(val);
                }
            })
            return newState;
        case LOAD_QUEUE:
            // console.log(`action.rows: ${JSON.stringify(action)}`)
            return action.rows;
        default:
            return state;
    }
}

export default combineReducers({
    queues,
});