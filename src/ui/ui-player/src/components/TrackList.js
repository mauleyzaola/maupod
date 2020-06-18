import React from 'react';
import {msToString, secondsToDate} from "../helpers";
import {Link} from "react-router-dom";
import {linkAlbumList, linkGenreList, linkPerformerList} from "../routes";

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

const TrackListRow = ({row}) => {
    row.recorded_date = row.recorded_date || '';
    const modifiedDate = row.modified_date ? secondsToDate(row.modified_date.seconds).toLocaleDateString() : '';
    const Thumbnail = () => {
        if(!row.sha_image){
            return null;
        }
        return <img style={{width: '30px'}}  alt='cover' src={`http://localhost:9000/thumbnail/${row.sha_image}.png`} />;
    }
    return (
        <tr>
            <td>
                <Thumbnail />
            </td>
            <td>{row.track_position}</td>
            <td>{row.track}</td>
            <td>
                <Link to={linkPerformerList(row)}>{row.performer}</Link>
            </td>
            <td>
                <Link to={linkGenreList(row)}>{row.genre}</Link>
            </td>
            <td>{msToString(row.duration)}</td>
            <td>
                <Link to={linkAlbumList(row)}>{row.album}</Link>
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