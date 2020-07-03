import React from 'react';
import { audioScan } from "./api";

class AudioScan extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            force: false,
        }
    }

    handleStartScan = () => {
        const { force } = this.state;
        audioScan({
            root:'/music-store',
            force,
        }).then(() => console.log('Request was successful'))
    }

    onChange = () => {
        this.setState({force: !this.state.force})
    }

    render() {
        return(
            <div>
                <p>Click the button to start audio scan</p>
                <input type="checkbox" checked={this.state.force} onChange={this.onChange} />
                <button type='button' onClick={this.handleStartScan}>Start Scan</button>
            </div>
        )
    }
}

export default AudioScan;