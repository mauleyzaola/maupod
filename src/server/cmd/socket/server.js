const WebSocket = require('ws')
const wsOptions = { port: 8080 };
const wss = new WebSocket.Server(wsOptions);
const NATS = require('nats');
const nc = NATS.connect({json: true});
const messages = require('./nodepb/messages_pb');
const ipcCommands = messages.IPCCommand;
const remoteCommands = messages.RemoteCommand;
const ipcMsg = messages.Message.MESSAGE_IPC;


console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

const sendPlay = media => {
    nc.publish(ipcMsg, {
        media: media,
        command: ipcCommands.IPC_PLAY,
    });
}

const sendPause = media => {
    nc.publish(ipcMsg, {
        media: media,
        command: ipcCommands.IPC_PAUSE,
    });
}

const sendVolume = ({media, value}) => {
    nc.publish(ipcMsg, {
        value,
        media,
        command: ipcCommands.IPC_VOLUME,
    });
}

wss.on('connection', ws => {
    const addr = ws._socket.remoteAddress
    console.log(`new connection from ${addr}`);

    ws.on('message', message => {
        const data = JSON.parse(message);
        try{
            switch (data.subject) {
                case remoteCommands.REMOTE_PLAY:
                    sendPlay(data.media);
                    break;
                case remoteCommands.REMOTE_PAUSE:
                    sendPause(data.media);
                    break;
                case remoteCommands.REMOTE_VOLUME:
                    sendVolume(data);
                    break;
                default:
                    console.log(`unsupported: ${data.subject}`);
                    break;
            }
        }catch (e) {
            console.log(e);
        }
    })

    ws.send('beer is good');
})

// working webcosket server
// TODO: connect to NATS and dispatch events to the UI