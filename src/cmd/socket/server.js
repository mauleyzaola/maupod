const WebSocket = require('ws')
const wsOptions = { port: 8080 };
const wss = new WebSocket.Server(wsOptions);
const servers = ['nats://nats-server:4222']
const NATS = require('nats');
const nc = NATS.connect({json: true, servers: servers});
const messages = require('./nodepb/messages_pb');
const subjects = messages.Message;

console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

nc.subscribe(subjects.MESSAGE_SOCKET_TRACK_POSITION_PERCENT, (msg) => {
    try{
        const { media, percent } = msg;
        console.log(`received event track: ${media.track} percent played: ${percent}`)
        const data = {
            subject:'MESSAGE_SOCKET_TRACK_POSITION_PERCENT',
            media,
            percent,
        }
        wss.clients.forEach(ws => {
            if(ws.isAlive === false) return ws.terminate();
            ws.send(JSON.stringify(data));
        })
    }catch (e){
        console.log(e);
    }
})

// TODO: need to call ws.send() passing the payload from NATS

wss.on('connection', ws => {
    const addr = ws._socket.remoteAddress
    console.log(`new connection from ${addr}`);
    // ws.on('message', message => {
    //     const data = JSON.parse(message);
    //     try{
    //         switch (data.subject) {
    //             case remoteCommands.REMOTE_PLAY:
    //                 sendPlay(data.media);
    //                 break;
    //             default:
    //                 console.log(`unsupported: ${data.subject}`);
    //                 break;
    //         }
    //     }catch (e) {
    //         console.log(e);
    //     }
    // })

    ws.send('socket started');
})
