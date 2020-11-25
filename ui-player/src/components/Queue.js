import React from 'react';
import { connect } from 'react-redux';
import {AlbumLink} from "./TrackList";
import {msToDate, msToString} from "../helpers";
import {Link} from "react-router-dom";
import {linkGenreList, linkPerformerList} from "../routes";
import {FaMinusSquare} from "react-icons/fa/index";
import {handleDeleteQueue, handleLoadQueue} from "../actions/queue";

const TrackListHeader = ({children}) => (
    <table className='table table-bordered table-hover table-striped'>
    <thead>
    <tr>
        <th>#</th>
        <th>Track</th>
        <th>Performer</th>
        <th>Genre</th>
        <th>Duration</th>
        <th>Album</th>
        <th>Year</th>
        <th>Format</th>
        <th></th>
    </tr>
    </thead>
        <tbody>
        {children}
        </tbody>
    </table>
)

const TrackListRow = ({index, row, onDelete}) => {
    row.recorded_date = row.recorded_date || '';
    return (
        <tr>
            <td>{index + 1}</td>
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
            <td>{row.recorded_date}</td>
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


class Queue extends React.Component{
    componentDidMount() {
        return this.loadData();
    }
    loadData = () => this.props.dispatch(handleLoadQueue());
    onDelete = index => this.props.dispatch(handleDeleteQueue({index}));

    render() {
        const { queues } = this.props;
        const totalMs = queues.reduce((x,i) => x + i.media.duration, 0)
        const toDate = msToDate(totalMs)
        const totalFormat = toDate.toTimeString().substring(0,8)
        return (
            <div>
                <h2>{`Queue List (${totalFormat})`}</h2>
                <TrackListHeader>
                    {queues.map((row, index) => <TrackListRow key={row.id} row={row.media} index={index} onDelete={this.onDelete} />)}
                </TrackListHeader>
            </div>
        )
    }
}

export default connect((state) => ({
    queues: state.queues
}))(Queue);
