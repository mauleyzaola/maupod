import React from 'react';
import { decodeURL, mediaList } from "./api";
import { msToString, secondsToDate } from "./helpers";

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
            <td>{row.performer}</td>
            <td>{row.genre}</td>
            <td>{msToString(row.duration)}</td>
            <td>{row.album}</td>
            <td>{row.sampling_rate}</td>
            <td>{row.recorded_date}</td>
            <td>{modifiedDate}</td>
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
        const uri = new URL(window.location.href);
        const search = decodeURL(uri);
        mediaList(search)
            .then(res => res.data || [])
            .then(rows => this.setState({rows}))
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