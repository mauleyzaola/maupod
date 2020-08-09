import React from 'react';
import {queueList} from "../api";
import {AlbumLink} from "./TrackList";
import {msToString, secondsToDate} from "../helpers";
import {Link} from "react-router-dom";
import {linkGenreList, linkPerformerList} from "../routes";

const TrackListHeader = () => (
    <thead>
    <tr>
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
    return (
        <tr>
            <td>{row.track_position}</td>
            <td>
                <span title={row.location}>
                    {row.track}
                </span>
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


// TODO: this list should be updated when a track is either added or removed from the queue
class Queue extends React.Component{
    state = {
        rows: [],
    }

    componentDidMount() {
        return this.loadData();
    }

    loadData = () => queueList().then(response => this.setState({rows: response.data.rows || []}));

    render() {
        const { rows } = this.state;
        return (
            <div>
                <h2>Queue List</h2>

                <table className='table table-bordered table-hover table-striped'>
                    <TrackListHeader />
                    <tbody>
                    {rows.map(row => <TrackListRow key={row.id} row={row} />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Queue;