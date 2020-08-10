import React from 'react';
import { connect } from 'react-redux';
import {AlbumLink} from "./TrackList";
import {msToString, secondsToDate} from "../helpers";
import {Link} from "react-router-dom";
import {linkGenreList, linkPerformerList} from "../routes";
import {FaMinusSquare} from "react-icons/fa/index";
import {handleDeleteQueue, handleLoadQueue} from "../actions/queue";

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
        <th></th>
    </tr>
    </thead>
)

const TrackListRow = ({index, row, onDelete}) => {
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
            <td>
                <span
                   title='remove track from queue'
                    onClick={() => onDelete(index)}
                >
                    <FaMinusSquare/>
                </span>
            </td>
        </tr>
    )
}


// TODO: this list should be updated when a track is either added or removed from the queue
class Queue extends React.Component{
    componentDidMount() {
        return this.loadData();
    }
    loadData = () => this.props.dispatch(handleLoadQueue());
    onDelete = index => this.props.dispatch(handleDeleteQueue({index}));

    render() {
        const { queues } = this.props;
        return (
            <div>
                <h2>Queue List</h2>

                <table className='table table-bordered table-hover table-striped'>
                    <TrackListHeader />
                    <tbody>
                    {queues.map((row, index) => <TrackListRow key={row.id} row={row.media} index={index} onDelete={this.onDelete} />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default connect((state) => ({
    queues: state.queues
}))(Queue);
