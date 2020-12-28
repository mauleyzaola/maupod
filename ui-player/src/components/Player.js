import React from 'react';
import PropTypes from 'prop-types';
import { FaForward, FaPlay, FaPause, FaAngleDoubleUp, FaAngleDoubleDown, FaVolumeDown, FaVolumeUp } from "react-icons/fa";
import {IPC_PAUSE, IPC_PLAY, IPC_SKIP, POSITION_BOTTOM, POSITION_TOP} from "../consts";
import API from "../api";

const TrackPlayControls = ({media}) => (
    <div className='form-inline'>
        <PlayerPlay media={media} />
        <PlayerPause media={media} />
        <PlayerSkip media={media} />
        <PlayerVolumeDown />
        <PlayerVolumeUp />
    </div>
)

const TrackListControls = ({media}) => (
    <div className='form-inline'>
        <PlayerPlay media={media} />
        <PlayerPlayNext media={media} />
        <PlayerPlayLater media={media} />
    </div>
)

class PlayerPlay extends React.Component{
    onClick = (media) => {
        API.ipcCommand(({ command: IPC_PLAY, media}))
            .then(() => {})
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
    onClick = (media) => {
        API.ipcCommand(({ command: IPC_PAUSE, media}))
            .then(() => {})
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
    onClick = media => {
        API.queueAdd({media: media, named_position: POSITION_TOP})
            .then(() => {})
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
    onClick = media => {
        API.queueAdd({media: media, named_position: POSITION_BOTTOM})
            .then(() => {})
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
    onClick = media => {
        API.ipcCommand(({ command: IPC_SKIP, media}))
            .then(() => {})
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

class PlayerVolumeDown extends React.Component{
    onClick = () => {
        API.volumeChange(({ offset: -5 }))
            .then(() => {})
    }

    render() {
        return (
            <button type="button"
                    title="volume down"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick()}>
                <FaVolumeDown />
            </button>
        )
    }
}

class PlayerVolumeUp extends React.Component{
    onClick = () => {
        API.volumeChange(({ offset: 5 }))
            .then(() => {})
    }

    render() {
        return (
            <button type="button"
                    title="volume up"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick()}>
                <FaVolumeUp />
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
    PlayerSkip,
    PlayerVolumeDown,
    PlayerVolumeUp,
    TrackListControls,
    TrackPlayControls,
}