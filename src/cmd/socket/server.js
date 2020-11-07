const WebSocket = require('ws')
const wsOptions = { port: `${process.env.MAUPOD_SOCKET_PORT}` };
const wss = new WebSocket.Server(wsOptions);
const servers = ['nats://nats-server:4222']
const NATS = require('nats');
const nc = NATS.connect({json: true, servers: servers});
const messages = {
    // ipc commands
    IPC_PLAY : 0,
    IPC_PAUSE : 1,
    IPC_LOAD : 2,
    IPC_VOLUME : 3,
    IPC_SKIP : 4, // not really IPC but in the same workflow

    MESSAGE_ARTWORK_SCAN : 10,
    MESSAGE_AUDIO_SCAN : 11,
    MESSAGE_MEDIA_INFO : 12,
    MESSAGE_MEDIA_UPDATE_ARTWORK : 13,
    MESSAGE_MEDIA_UPDATE_SHA : 14,
    MESSAGE_MEDIA_DELETE : 15,
    MESSAGE_TAG_UPDATE : 16,
    MESSAGE_MEDIA_UPDATE : 17,
    MESSAGE_MEDIA_SPECTRUM_GENERATE : 18,
    MESSAGE_MEDIA_DB_SELECT : 19,
    MESSAGE_MEDIA_EXTRACT_ARTWORK_FROM_FILE : 20,
    MESSAGE_UPSERT_MEDIA_EVENT : 21,
    MESSAGE_ARTWORK_FIND_ALBUM_COVER : 22,
    MESSAGE_ARTWORK_DOWNLOAD : 23,
    MESSAGE_SYNC_FILES : 24,
    MESSAGE_SHA_SCAN : 25,

    MESSAGE_IPC : 100,

    // specific to mpv events
    MESSAGE_MPV_EOF_REACHED : 202,
    MESSAGE_MPV_PERCENT_POS : 203,
    MESSAGE_MPV_TIME_POS : 204,
    MESSAGE_MPV_TIME_REMAINING : 205,

    // specific to player events
    MESSAGE_EVENT_ON_TRACK_STARTED : 250,
    MESSAGE_EVENT_ON_TRACK_FINISHED : 251,
    MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE : 252,
    MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE : 253,

    // queue management
    MESSAGE_QUEUE_LIST : 300,
    MESSAGE_QUEUE_ADD : 301,
    MESSAGE_QUEUE_REMOVE : 302,

    // file management
    MESSAGE_DIRECTORY_READ : 400,

    // micro service discovery
    MESSAGE_MICRO_SERVICE_ARTWORK : 420,
    MESSAGE_MICRO_SERVICE_AUDIOSCAN : 421,
    MESSAGE_MICRO_SERVICE_MEDIAINFO : 422,
    MESSAGE_MICRO_SERVICE_PLAYER : 423,
    MESSAGE_MICRO_SERVICE_RESTAPI : 424,
    MESSAGE_MICRO_SERVICE_SOCKET : 425,

    // real time socket events
    MESSAGE_SOCKET_TRACK_POSITION_PERCENT : 500,
    MESSAGE_SOCKET_TRACK_POSITION_PERCENT_CHANGE : 501,
    MESSAGE_SOCKET_QUEUE_CHANGE : 502,
    MESSAGE_SOCKET_PLAY_TRACK : 503,
}

// TODO: import from src/proto directory and generate the JS files there

console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

const broadcastMessage = data => {
    try{
        wss.clients.forEach(ws => {
            if(ws.isAlive === false) return ws.terminate();
            ws.send(JSON.stringify(data));
        })
    }catch (e){
        console.log(e);
    }
}

nc.subscribe(messages.MESSAGE_SOCKET_TRACK_POSITION_PERCENT, (msg) => {
    const { media, percent, seconds, secondsTotal } = msg;
    const data = {
        subject:'MESSAGE_SOCKET_TRACK_POSITION_PERCENT',
        percent,
        seconds,
        seconds_total: secondsTotal,
        media,
    }
    broadcastMessage(data);
})

nc.subscribe(messages.MESSAGE_SOCKET_QUEUE_CHANGE, (msg) => {
    broadcastMessage({
        subject: 'MESSAGE_SOCKET_QUEUE_CHANGE',
    });
})

nc.subscribe(messages.MESSAGE_SOCKET_PLAY_TRACK, (msg) => {
    const { media } = msg;
    broadcastMessage({
        subject: 'MESSAGE_SOCKET_PLAY_TRACK',
        media,
    });
})


wss.on('connection', ws => {
    const addr = ws._socket.remoteAddress
    console.log(`new connection from ${addr}`);

    ws.on('message', message => {
        try {
            const data = JSON.parse(message);
            switch (data.subject){
                // triggered from the front end when user changes track position
                case 'MESSAGE_SOCKET_TRACK_POSITION_PERCENT':
                    const { media, percent } = data;
                    const payload = {
                        percent,
                        media,
                    }
                    nc.publish(messages.MESSAGE_SOCKET_TRACK_POSITION_PERCENT_CHANGE, payload);
                    break;
                default:
                    break;
            }
        }catch (e){
            console.log(e);
        }
    })
})
