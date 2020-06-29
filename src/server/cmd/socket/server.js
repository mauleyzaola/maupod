const WebSocket = require('ws')
const wsOptions = { port: 8080 };
const wss = new WebSocket.Server(wsOptions);
const NATS = require('nats');
const nc = NATS.connect();
const protos = require('google-proto-files');

async function loadProtobuf(){
    const files = protos.getProtoPath('logging', 'v2');
    console.log(files);

    const root = await protos.load(__dirname + './../../proto/messages.proto');
    console.log(root);
}

loadProtobuf();

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