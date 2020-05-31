import React from 'react';
import { audioScan } from "./api";

class AudioScan extends React.Component{
    constructor(props) {
        super(props);
        this.handleStartScan.bind = this.handleStartScan;
    }

    handleStartScan = () => {
        audioScan({
            root:'/music-store',
        }).then(() => alert('Request was successful'))
    }

    render() {
        return(
            <div>
                <p>Click the button to start audio scan</p>
                <button type='button' onClick={this.handleStartScan}>Start Scan</button>
            </div>
        )
    }
}

export default AudioScan;