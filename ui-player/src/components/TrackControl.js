import React from 'react';
import { connect } from 'react-redux';
import {TrackPlayControls} from "./Player";
import {handleLoadQueue} from "../actions/queue";
import {applyBlur, CANVAS_WIDTH, loadCanvasImage} from "../canvas";

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
                    case 'MESSAGE_SOCKET_QUEUE_CHANGE':
                        this.props.dispatch(handleLoadQueue());
                        break;
                    case 'MESSAGE_SOCKET_PLAY_TRACK':
                        this.drawTrackSpectrum(data.media);
                        break;
                    default:
                        break;
                }
            }catch (e){
                console.warn(e);
            }
        });
        window.addEventListener('resize', () => this.setState({width: this.windowSize()}));
        this.setState({width: this.windowSize()});
    }

    secondsToDisplay = seconds => {
        const t = new Date(seconds * 1000);
        const secs = t.getSeconds();
        const mins = t.getMinutes();
        const offsetSecs = secs < 10 ? '0' : '';
        const offsetMin = mins < 10 ? '0' : '';
        return `${offsetMin}${mins}:${offsetSecs}${secs}`;
    }

    drawTrackSpectrum = media => {
        const canvas = document.getElementById('canvas');
        loadCanvasImage({
            canvas,
            src: `${process.env.REACT_APP_MAUPOD_API}/media/${media.id}/spectrum`,
        })
    }

    onMessageReceived = data => {
        const { media } = this.state;
        const { percent, seconds, seconds_total } = data;
        this.setState({
            percent: data.percent,
            media: data.media,
            timePlayed: this.secondsToDisplay(seconds),
            timeTotal: this.secondsToDisplay(seconds_total),
        });
        if(data.media.id !== media.id){
            this.drawTrackSpectrum(data.media);
        } else{
            // TODO: handle play song from the start, probably another message?
            const x1 = 0;
            const x2 = CANVAS_WIDTH * percent / 100;
            applyBlur({x1,x2});
        }
    }

    onPositionChange = e => {
        const {  media } = this.state;
        if(!media.id) return null;
        let percent = parseFloat(e.target.value);
        const data = {
            subject: 'MESSAGE_SOCKET_TRACK_POSITION_PERCENT',
            media,
            percent,
        }
        this.ws.send(JSON.stringify(data));
    }

    // TODO: allow to set the position in the spectrum

    render() {
        const { media, timePlayed, timeTotal } = this.state;
        // const { width } = this.state;
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
                    <div id='spectrum_div'>
                    </div>
                    <div>
                        <canvas id='canvas'/>
                    </div>
                </div>
            </div>
        )
    }
}

// TODO: handle external events using redux, for now props are empty and all is handled locally
export default connect((state) => ({ }))(TrackControl);
