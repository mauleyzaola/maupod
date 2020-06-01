import React from 'react';
import { Link } from "react-router-dom";
import { decodeURL, mediaList } from "./api";
import { msToString, secondsToDate } from "./helpers";
import {linkAlbumList, linkGenreList, linkPerformerList} from "./routes";

const MediaHeader = () => (
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

const MediaLine = ({row}) => {
    row.recorded_date = row.recorded_date || '';
    const modifiedDate = row.modified_date ? secondsToDate(row.modified_date.seconds).toLocaleDateString() : '';
    return (
        <tr>
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

class MediaList extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        const uri = decodeURL(window.location.search);
        this.loadData(uri);
    }

    loadData = search => {
        mediaList(search)
            .then(res => res.data || [])
            .then(rows => this.setState({rows}))
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if(JSON.stringify(prevProps.location) === JSON.stringify(this.props.location)){
            return;
        }
        this.loadData(decodeURL(this.props.location.search));
    }

    render() {
        const { rows } = this.state;
        return (
            <div>
                <table>
                    <MediaHeader />
                    <tbody>
                    {rows.map(row => <MediaLine key={row.id} row={row} />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default MediaList;