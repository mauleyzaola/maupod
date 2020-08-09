import React from 'react';
import {TrackPlayControls} from "./Player";
import {spectrumImage} from "../api";

class TrackControl extends React.Component{
    state = {
        progress: 0,
        media: {},
        width: 0,
        timePlayed: '',
        timeTotal: '',
    }
    ws;

    windowSize = () => window.innerWidth - 50;

    componentDidMount() {
        this.ws = new WebSocket(`${process.env.REACT_APP_MAUPOD_SOCKET}`);
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
        window.addEventListener('resize', () => this.setState({width: this.windowSize()}));
        this.setState({width: this.windowSize()});
    }

    secondsToDisplay = seconds => {
        if(!seconds) return '';
        const t = new Date(seconds * 1000);
        const secs = t.getSeconds();
        const mins = t.getMinutes();
        const offsetSecs = secs < 10 ? '0' : '';
        const offsetMin = mins < 10 ? '0' : '';
        return `${offsetMin}${mins}:${offsetSecs}${secs}`;
    }

    onMessageReceived = data => {
        this.setState({
            percent: data.percent,
            media: data.media,
            timePlayed: this.secondsToDisplay(data.seconds),
            timeTotal: this.secondsToDisplay(data.seconds_total),
        });
    }

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

    // TODO: allow to set the position in the spectrum
    // <input type='range' className='form-control' min='0' max='100' value={percent} onChange={this.onPositionChange} />

    render() {
        const { media, percent, width, timePlayed, timeTotal } = this.state;
        if(!media || !media.id) return null;
        return (
            <div className='row'>
                <div className='col'>
                    <div>
                        <strong>{timePlayed} / {timeTotal} </strong>
                        {media.performer} |
                        {media.album} |
                        {media.track}
                    </div>
                    <TrackPlayControls media={media} />
                    {media.id ? <img src={`${process.env.REACT_APP_MAUPOD_API}/media/${media.id}/spectrum`} width={`${width}px`} height="150px" /> : null}
                </div>
            </div>
        )
    }
}

export default TrackControl;