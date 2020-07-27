import React from 'react';
import PropTypes from 'prop-types';
import { FaPlay, FaAngleDoubleUp, FaAngleDoubleDown } from "react-icons/all";
import {IPC_PLAY, POSITION_BOTTOM, POSITION_TOP} from "../consts";
import {ipcCommand, queueAdd} from "../api";


const cleanMedia = media => {
    const result = Object.assign({}, media);
    result.recorded_date = result.recorded_date || 0;
    return result;
}


const TrackPlayControls = ({media}) => (
    <div className='form-inline'>
        <PlayerPlay media={media} />
        <PlayerPlayNext media={media} />
        <PlayerPlayLater media={media} />
    </div>
)


class PlayerPlay extends React.Component{
    constructor(props) {
        super(props);
    }

    onClick = (media) => {
        ipcCommand(({ command: IPC_PLAY, media: cleanMedia(media)}))
            .then(data => console.log(data))
    }

    render() {
        const { media } = this.props;
        return (
            <button type="button"
                    className="btn btn-secondary btn-sm"
                    onClick={() => this.onClick(media)}>
                <FaPlay />
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

PlayerPlay.propTypes = {
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
    TrackPlayControls,
}