import React from 'react';
import {albumViewList, decodeURL} from "./api";
import {linkAlbumSongList} from "./routes";
import { Link } from "react-router-dom";
import {msToString} from "./helpers";

const AlbumCard = ({r}) => {
    return (
        <div className='album-card col-3'>
            <div className="card text-white bg-primary">
                <div className="card-header">
                    <Link to={linkAlbumSongList(r)}>
                    {r.album}
                    </Link>
                </div>
                <div className="card-body">
                    <h4 className="card-title">{r.performer}</h4>
                    <p className="card-text">
                        Genre: {r.genre}
                    </p>
                    <p className="card-text">
                        {r.recorded_date ? `Recorded Date: ${r.recorded_date}` : null}
                    </p>
                    <p className="card-text">
                        {r.track_name_total ? `Track Count: ${r.track_name_total}` : null}
                    </p>
                    <p className="card-text">
                        {r.format ? `Format: ${r.format}` : null}
                    </p>
                </div>
                <div className="card-footer">
                    <small className="text-muted">
                        {r.duration ? `Duration: ${msToString(r.duration)}` : null}
                    </small>
                </div>
            </div>
        </div>
    )
}

                // <Link to={linkAlbumList(row)}>

class Albums extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        const data = decodeURL(this.props.location.search);
        // data.limit=50;
        albumViewList(data).then(res => res.data || [])
           .then(rows => this.setState({rows}));
    }

    render() {
        const { rows } = this.state;
        return(
            <div className='card-deck'>
                {rows.map(r => <AlbumCard key={r.id} r={r} />)}
            </div>
        )
    }
}

export default Albums;