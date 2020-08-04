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

    ws.on('message', message => {
        try {
            const data = JSON.parse(message);
            switch (data.subject){
                case 'MESSAGE_SOCKET_TRACK_POSITION_PERCENT':
                    const { media, percent } = data;
                    const payload = {
                        media,
                        percent,
                    }
                    nc.publish(subjects.MESSAGE_SOCKET_TRACK_POSITION_PERCENT_CHANGE, payload);
                    break;
                default:
                    break;
            }
        }catch (e){
            console.log(e);
        }
    })
})
