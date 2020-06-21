import {ipcCommand} from "./api";
import {IPC_LOAD, IPC_PLAY} from "./consts";

const playTrack = r => {
    // TODO: fix this hack with the recorded_date being empty string
    if(r.recorded_date === ''){
        r.recorded_date = 0;
    }
    console.log(`requested to play file: ${r.location}`);
    const promises = [
        ipcCommand({
            command: IPC_LOAD,
            media: r,
        }),
        ipcCommand({
            command: IPC_PLAY,
            media: r,
        }),
    ]
    Promise.all(promises).then(() => {});
}

export {
    playTrack,
}