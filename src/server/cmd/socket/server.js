const WebSocket = require('ws')
const wsOptions = { port: 8080 };
const wss = new WebSocket.Server(wsOptions);
const NATS = require('nats');
const nc = NATS.connect();
const messages = require( './../../nodepb/pb');

// console.log(`proto.pb.Message.MESSAGE_ARTWORK_SCAN: ${messages}`);

console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

wss.on('connection', ws => {
    const addr = ws._socket.remoteAddress
    console.log(`new connection from ${addr}`);
    ws.on('message', message => {
        console.log(`Received message => ${message}`)
    })
    ws.send('ho!')
})

// working webcosket server
// TODO: connect to NATS and dispatch events to the UI