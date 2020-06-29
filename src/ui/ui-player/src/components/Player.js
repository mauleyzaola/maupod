import React from 'react';
import PropTypes from 'prop-types';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { FaPlay, FaPause } from "react-icons/all";
import {REMOTE_PAUSE, REMOTE_PLAY} from "../consts";

const socket = new W3CWebSocket(`ws://localhost:8080`);

// missing fields, issues with deserializing data in the server
const cleanMedia = media => {
    const {id, track, location} = media;
    return {id, track, location};
}

const sendWSMessage = data => socket.send(JSON.stringify(data));

class Player extends React.Component{
    // constructor(props) {
    //     super(props);
    // }

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
        console.log("websocket will close");
        socket.close();
    }

    onPause = (media) => {
        // we need to send data as string
        // ideally we should use protobuf all over the places
        // this is the current workflow
        // browser -> sends JSON -> nodejs -> parses JSON -> creates protobuf message -> sends to NATS

        sendWSMessage({
            subject: REMOTE_PAUSE,
            media: cleanMedia(media),
        });
    }

    onPlay = (media) => sendWSMessage({ subject: REMOTE_PLAY, media: cleanMedia(media)});


    render() {
        const { visible } = this.props;
        if(!visible){
            return null;
        }
        return (
            <ul className='pagination'>
                <li className='page-item'>
                    <button type='button' className='btn btn-secondary btn-sm' onClick={() => this.onPlay(this.props.media)}>
                        <FaPlay />
                    </button>
                </li>
                <li className='page-item'>
                    <button type='button' className='btn btn-secondary btn-sm' onClick={() => this.onPause(this.props.media)}>
                        <FaPause />
                    </button>
                </li>
            </ul>
        )
    }
}

Player.propTypes = {
    visible: PropTypes.bool.isRequired,
    media: PropTypes.object.isRequired,
}

export default Player;