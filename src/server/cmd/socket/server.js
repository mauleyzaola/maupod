const WebSocket = require('ws')
const wsOptions = { port: 8080 };
const wss = new WebSocket.Server(wsOptions);
const NATS = require('nats');
const nc = NATS.connect();
const messages = require('./nodepb/messages_pb');
const remoteMessages = messages.RemoteCommand;
const ipcMessages = messages.IPCCommand;

console.log(JSON.stringify(remoteMessages))
console.log(JSON.stringify(ipcMessages))

console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

wss.on('connection', ws => {
    const addr = ws._socket.remoteAddress
    console.log(`new connection from ${addr}`);

    ws.on('message', message => {
        const data = JSON.parse(message);
        console.log(`${JSON.stringify(data)}`);
    })

    ws.send('ho!')
})

// working webcosket server
// TODO: connect to NATS and dispatch events to the UI