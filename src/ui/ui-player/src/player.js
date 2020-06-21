import {ipcCommand} from "./api";
import {IPC_LOAD, IPC_PLAY} from "./consts";

const playTrack = r => {
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