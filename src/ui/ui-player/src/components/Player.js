import React from 'react';
import {bool} from 'prop-types';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { FaPlay, FaPause } from "react-icons/all";

const client = new W3CWebSocket(`ws://localhost:8080`);

class Player extends React.Component{
    // constructor(props) {
    //     super(props);
    // }

    // stupid simple connection which is working
    componentDidMount() {
        client.onopen = () => {
            console.log('websocket connected')
        }
        client.onmessage = (message) => {
            const { data } = message;
            console.log(data);
        };
    }

    componentWillUnmount() {
        console.log("websocket will close");
        client.close();
    }


    render() {
        const { visible } = this.props;
        if(!visible){
            return null;
        }
        return (
            <div className='row'>
                <div className='col'>
                    <div className='btn-toolbar'>
                        <div className='btn-group mr-2'>
                            <button type='button' className='btn btn-secondary'>
                                <FaPlay />
                            </button>
                            <button type='button' className='btn btn-secondary'>
                                <FaPause />
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

Player.propTypes = {
    visible: bool.isRequired,
}

export default Player;