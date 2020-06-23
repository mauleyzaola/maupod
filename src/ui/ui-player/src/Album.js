import React from 'react';
import AlbumHeader from "./components/AlbumHeader";
import {decodeURL} from "./api";
import { albumViewList, mediaList } from "./api";
import {msToString } from "./helpers";
import {playTrack} from "./player";


const TrackListHeader = () => (
    <thead>
    <tr>
        <th>#</th>
        <th>Track</th>
        <th>Duration</th>
        <th>Format</th>
    </tr>
    </thead>
)

const TrackListRow = ({row, onTrackClick}) => {
    return (
        <tr>
            <td>{row.track_position}</td>
            <td onClick={() => onTrackClick(row)}>{row.track}</td>
            <td>{msToString(row.duration)}</td>
            <td>{row.format}</td>
        </tr>
    )
}


class Album extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            album:null,
            rows: [],
            genre: '',
        }
    }

    loadData = search => {
        let album = null;
        albumViewList(search)
            .then(response => {
                const data = response.data || [];
                if(data.length !== 1) return;
                album  = data[0];
            })
            .then(() => mediaList({ sort:'track_position', direction: 'asc', ...search}))
            .then(res => res.data || [])
            .then(rows => this.setState({rows, album}))
    }

    componentDidMount() {
        const uri = decodeURL(window.location.search);
        this.loadData(uri);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        const { rows } = this.state;
        if(prevState.rows.length !== rows.length && rows.length !== 0){
            const genre = rows[0].genre;
            this.setState({genre});
        }
    }

    onTrackClick = (r) => playTrack(r);

    render() {
        const { album, rows } = this.state;
        return (
            <div>
                <AlbumHeader album={album} />
                <table className='table table-bordered table-hover table-striped'>
                    <TrackListHeader />
                    <tbody>
                    {rows.map(row => <TrackListRow key={row.id} row={row} onTrackClick={this.onTrackClick} />)}
                    </tbody>
                </table>
            </div>
        );
    }
}

export default Album;