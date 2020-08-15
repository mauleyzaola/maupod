import React from 'react';
import { connect } from 'react-redux';
import {TrackPlayControls} from "./Player";
import {handleLoadQueue} from "../actions/queue";

class TrackControl extends React.Component{
    state = {
        progress: 0,
        media: {},
        width: 0,
        timePlayed: '',
        timeTotal: '',
    }
    ws;
    ctx;
    imageData;

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

    onMessageReceived = data => {
        const { media } = this.state;
        this.setState({
            percent: data.percent,
            media: data.media,
            timePlayed: this.secondsToDisplay(data.seconds),
            timeTotal: this.secondsToDisplay(data.seconds_total),
        });
        if(data.media.id !== media.id){
            const img = new Image();
            img.crossOrigin = '';
            img.src=`${process.env.REACT_APP_MAUPOD_API}/media/${data.media.id}/spectrum`
            const canvas = document.getElementById('canvas');
            canvas.width = 1920;
            this.ctx = canvas.getContext('2d');
            img.addEventListener('load', () => {
                this.ctx.drawImage(img, 0, 0, 1920, 150);
                const image = this.ctx.getImageData(0,0,1920/2,150);
                this.imageData = image.data;
                // const length = this.imageData.length;
                // for(let i = 0; i < length; i += 4){
                //     const color = this.imageData[i].toString();
                // }
                // this.ctx.putImageData(image,0,0);
            })
        }else if(this.imageData && false){
            // for(let i = 3; i < length; i += 4){
                // this.imageData.data[i] = 50;
            // }
            // this.ctx.putImageData(this.imageData,0,0);
        }
    }
    /*
    declare var ImageData: {
    prototype: ImageData;
    new(width: number, height: number): ImageData;
    new(array: Uint8ClampedArray, width: number, height?: number): ImageData;
};

     */

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
    // <input type='range' className='form-control' min='0' max='100' value={percent} onChange={this.onPositionChange} />

    render() {
        const { media, timePlayed, timeTotal } = this.state;
        const { width } = this.state;
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
