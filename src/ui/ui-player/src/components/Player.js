import React from 'react';
import {bool} from 'prop-types';
import { w3cwebsocket as W3CWebSocket } from "websocket";
const client = new W3CWebSocket(`ws://localhost:8080`);


class Player extends React.Component{
    constructor(props) {
        super(props);
    }

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


    render() {
        const { visible } = this.props;
        if(!visible){
            return null;
        }
        return (
            <div className='row'>
                <div className='col'>
                    TODO: display the play control here
                </div>
            </div>
        )
    }
}

Player.propTypes = {
    visible: bool.isRequired,
}

export default Player;