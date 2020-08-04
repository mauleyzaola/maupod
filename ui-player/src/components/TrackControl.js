import React from 'react';
import {PlayerPause, PlayerPlay, PlayerPlayNext, PlayerSkip, TrackPlayControls} from "./Player";

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

    onPositionChange = e => {
        const {  media } = this.state;
        if(!media.id) return null;
        let percent = parseFloat(e.target.value);
        if(percent <0 || percent > 100){
            console.warn(`percent out of range: ${percent}`)
            return;
        }
        const data = {
            subject: 'MESSAGE_SOCKET_TRACK_POSITION_PERCENT',
            media,
            percent,
        }
        this.ws.send(JSON.stringify(data));
    }

    render() {
        const { percent, media } = this.state;
        if(!media || !media.id) return null;
        return (
            <div className='row'>
                <div className='col'>
                    <div>
                        <strong>Performer: </strong>{media.performer} |
                        <strong>Album: </strong>{media.album} |
                        <strong>Track: </strong>{media.track}
                    </div>
                    <TrackPlayControls media={media} />
                    <input type='range' className='form-control' min='0' max='100' value={percent} onChange={this.onPositionChange} />
                </div>
            </div>
        )
    }
}

export default TrackControl;