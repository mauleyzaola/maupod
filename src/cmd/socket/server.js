const WebSocket = require('ws')
const wsOptions = { port: `${process.env.MAUPOD_SOCKET_PORT}` };
const wss = new WebSocket.Server(wsOptions);
const servers = ['nats://nats-server:4222']
const NATS = require('nats');
const nc = NATS.connect({json: true, servers: servers});
const messages = require('./nodepb/messages_pb');
const subjects = messages.Message;

console.log(`started websocket server on: ${JSON.stringify(wsOptions)}`);

nc.subscribe(subjects.MESSAGE_SOCKET_TRACK_POSITION_PERCENT, (msg) => {
    try{
        const { media, percent, seconds, secondsTotal } = msg;
        const data = {
            subject:'MESSAGE_SOCKET_TRACK_POSITION_PERCENT',
            percent,
            seconds,
            seconds_total: secondsTotal,
            media,
        }
        wss.clients.forEach(ws => {
            if(ws.isAlive === false) return ws.terminate();
            ws.send(JSON.stringify(data));
        })
    }catch (e){
        console.log(e);
    }
})

nc.subscribe(subjects.MESSAGE_SOCKET_QUEUE_CHANGE, (msg) => {
  try {
      wss.clients.forEach(ws => {
          if(ws.isAlive === false) return ws.terminate();
          ws.send(JSON.stringify({
              subject: 'MESSAGE_SOCKET_QUEUE_CHANGE',
          }));
      })
  }  catch (e){
      console.log(e);
  }
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
