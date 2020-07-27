import React from 'react';
import PropTypes from 'prop-types';
import { FaForward, FaPlay, FaPause, FaAngleDoubleUp, FaAngleDoubleDown } from "react-icons/all";
import {IPC_PAUSE, IPC_PLAY, IPC_SKIP, POSITION_BOTTOM, POSITION_TOP} from "../consts";
import {ipcCommand, queueAdd} from "../api";

const TrackPlayControls = ({media}) => (
    <div className='form-inline'>
        <PlayerPlay media={media} />
        <PlayerPause media={media} />
        <PlayerPlayNext media={media} />
        <PlayerPlayLater media={media} />
        <PlayerSkip media={media} />
    </div>
)

class PlayerPlay extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = (media) => {
        ipcCommand(({ command: IPC_PLAY, media}))
            .then(data => console.log(data))
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    title="play"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaPlay />
            </button>
        )
    }
}

class PlayerPause extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = (media) => {
        ipcCommand(({ command: IPC_PAUSE, media}))
            .then(data => console.log(data))
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    title="pause"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaPause />
            </button>
        )
    }
}

class PlayerPlayNext extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = media => {
        queueAdd({media: media, named_position: POSITION_TOP})
            .then(data => console.log(data));
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    title="play next"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaAngleDoubleUp />
            </button>
        )
    }
}

class PlayerPlayLater extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = media => {
        queueAdd({media: media, named_position: POSITION_BOTTOM})
            .then(data => console.log(data));
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    title="play later"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaAngleDoubleDown />
            </button>
        )
    }
}

class PlayerSkip extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = media => {
        ipcCommand(({ command: IPC_SKIP, media}))
            .then(data => console.log(data))
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    title="skip"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaForward />
            </button>
        )
    }
}

PlayerPlay.propTypes = {
    media: PropTypes.object.isRequired,
}

PlayerPause.propTypes = {
    media: PropTypes.object.isRequired,
}

PlayerSkip.propTypes = {
    media: PropTypes.object.isRequired,
}

PlayerPlayNext.propTypes = {
    media: PropTypes.object.isRequired,
}

TrackPlayControls.propTypes = {
    media: PropTypes.object.isRequired,
}

export {
    PlayerPlay,
    PlayerPlayNext,
    PlayerPause,
    TrackPlayControls,
}