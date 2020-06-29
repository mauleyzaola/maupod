import React from 'react';
import PropTypes from 'prop-types';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { FaPlay, FaPause } from "react-icons/all";
import {REMOTE_PAUSE, REMOTE_PLAY, REMOTE_VOLUME} from "../consts";

const socket = new W3CWebSocket(`ws://localhost:8080`);

// missing fields, issues with deserializing data in the server
const cleanMedia = media => {
    const result = Object.assign({}, media);
    result.recorded_date = result.recorded_date || 0;
    return result;
}

const sendWSMessage = data => socket.send(JSON.stringify(data));

class Player extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            volume: 100,
        }
    }

    currentMedia;

    // stupid simple connection which is working
    componentDidMount() {
        socket.onopen = () => {
            console.log('websocket connected')
        }
        socket.onmessage = (message) => {
            const { data } = message;
            console.log(data);
        };
    }

    componentWillUnmount() {
        // console.log("websocket will close");
        // socket.close();
    }

    onPause = (media) => {
        // we need to send data as string
        // ideally we should use protobuf all over the places
        // this is the current workflow
        // browser -> sends JSON -> nodejs -> parses JSON -> creates message -> sends to NATS as JSON

        sendWSMessage({
            subject: REMOTE_PAUSE,
            media: cleanMedia(media),
        });
    }

    onPlay = (media) => {
        this.currentMedia = media;
        sendWSMessage({ subject: REMOTE_PLAY, media: cleanMedia(media)});
    }

    onVolumeChange = e => {
        const volume = e.target.value || '0'
        this.setState({volume});
        if(!this.currentMedia) return;
        sendWSMessage({ subject: REMOTE_VOLUME, media: cleanMedia(this.currentMedia), value: volume });
    }


    render() {
        const { visible } = this.props;
        if(!visible){
            return null;
        }
        return (
            <div className='form-inline'>
                <button type='button' className='btn btn-secondary btn-sm' onClick={() => this.onPlay(this.props.media)}>
                    <FaPlay />
                </button>
                <button type='button' className='btn btn-secondary btn-sm' onClick={() => this.onPause(this.props.media)}>
                    <FaPause />
                </button>
                <input type='range' className='form-cotrol' min='0' max='130' value={this.state.volume} onChange={this.onVolumeChange} />
            </div>
        )
    }
}

Player.propTypes = {
    visible: PropTypes.bool.isRequired,
    media: PropTypes.object.isRequired,
}

export default Player;