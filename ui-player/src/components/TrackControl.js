import React from 'react';
import {ProgressBar} from "react-bootstrap";

class TrackControl extends React.Component{
    state = {
        progress: 0,
        media: {},
    }
    ws;
    componentDidMount() {
        this.ws = new WebSocket('ws://localhost:8080');
        this.ws.addEventListener('message', e => {
            try{
                const data = JSON.parse(e.data);
                switch (data.subject){
                    case 'MESSAGE_SOCKET_TRACK_POSITION_PERCENT':
                        this.onMessageReceived(data);
                        break;
                    default:
                        break;
                }
            }catch (e){}
        });
    }

    onMessageReceived = data => this.setState({percent: data.percent, media: data.media})

    render() {
        const { percent, media } = this.state;
        if(!media.id) return null;
        return (
            <div className='row'>
                <div className='col'>
                    <div>
                        <strong>Performer: </strong>{media.performer} |
                        <strong>Album: </strong>{media.album} |
                        <strong>Track: </strong>{media.track}
                    </div>
                    <ProgressBar now={percent} variant='info' />
                </div>
            </div>
        )
    }
}

export default TrackControl;