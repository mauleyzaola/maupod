import React from 'react';
import {msToString, secondsToDate} from "../helpers";
import {Link} from "react-router-dom";
import { linkAlbumView, linkGenreList, linkPerformerList } from "../routes";
import {TrackListControls} from "./Player";

const TrackListHeader = () => (
    <thead>
    <tr>
        <th></th>
        <th>#</th>
        <th>Track</th>
        <th>Performer</th>
        <th>Genre</th>
        <th>Duration</th>
        <th>Album</th>
        <th>Sampling Rate</th>
        <th>Year</th>
        <th>Last Modified</th>
        <th>Format</th>
    </tr>
    </thead>
)

const AlbumLink = ({r}) => r.album_identifier ? <Link to={linkAlbumView(r)}>{r.album}</Link> : <span>{r.album}</span>;

const TrackListRow = ({row}) => {
    row.recorded_date = row.recorded_date || '';
    const modifiedDate = row.modified_date ? secondsToDate(row.modified_date.seconds).toLocaleDateString() : '';
    const Thumbnail = () => {
        if(!row.image_location){
            return null;
        }
        return <img style={{width: '30px'}}  alt='cover' src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/${row.image_location}`} />;
    }
    return (
        <tr>
            <td>
                <Thumbnail />
            </td>
            <td>{row.track_position}</td>
            <td>
                <div className='row'>
                    <div className='col-4'>
                        <TrackListControls media={row}/>
                    </div>
                    <div className='col-8'>
                        <span title={row.location}>
                            {row.track}
                        </span>
                    </div>
                </div>
            </td>
            <td>
                <Link to={linkPerformerList(row)}>{row.performer}</Link>
            </td>
            <td>
                <Link to={linkGenreList(row)}>{row.genre}</Link>
            </td>
            <td>{msToString(row.duration)}</td>
            <td>
                <AlbumLink r={row} />
            </td>
            <td>{row.sampling_rate}</td>
            <td>{row.recorded_date}</td>
            <td>{modifiedDate}</td>
            <td>{row.format}</td>
        </tr>
    )
}

export {
    TrackListHeader,
    TrackListRow,
}