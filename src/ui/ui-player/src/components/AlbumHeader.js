import React from 'react';
import {msToString} from "../helpers";
import { Link } from "react-router-dom";
import { linkGenreList, linkPerformerList } from "../routes";

const Thumbnail = ({album}) => {
    if(!album.sha_image){
        return null;
    }
    return (
        <img alt='cover' src={`http://localhost:9000/${album.sha_image}.png`} />
    )
}


const AlbumHeader = ({album}) => {
    if(!album) return null;
    return(
        <div className='row'>
            <div className='col'>
                <div className='text-left'>
                    <Thumbnail album={album} />
                </div>
            </div>
            <div className='col'>
                <p>
                    <strong>Album: </strong>
                    {album.album}
                </p>
                <p>
                    <strong>Performer: </strong>
                    <Link to={linkPerformerList(album)}>
                        {album.performer}
                    </Link>
                </p>
                <p>
                    <strong>Genre: </strong>
                    <Link to={linkGenreList(album)}>
                        {album.genre}
                    </Link>
                </p>
                <p>
                    <strong>Duration: </strong>
                    {album.duration ? `${msToString(album.duration)}` : null}
                </p>
                <p>
                    <strong>Year: </strong>
                    {album.recorded_date}
                </p>
                <p>
                    <strong>Sampling Rate: </strong>
                    {album.sampling_rate ? `${album.sampling_rate}` : null}
                </p>
                <p>
                    <strong>Format: </strong>
                    {album.format}
                </p>
            </div>
        </div>
    )
}

export default AlbumHeader;