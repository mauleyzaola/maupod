import React from 'react';
import {albumViewList, decodeURL} from "./api";
import {linkAlbumView, linkGenreList, linkPerformerList} from "./routes";
import { Link } from "react-router-dom";
import {msToString} from "./helpers";


const Thumbnail = ({album}) => {
    if(!album.image_location){
        return null;
    }
    return (
        <img className='artwork-small' alt='cover' src={`http://localhost:9000/${album.image_location}`} />
    )
}


const AlbumCard = ({r}) => {
    return (
        <div className='album-card col-3'>
            <div className="card text-white bg-primary">
                <div className="card-header">
                    <Link to={linkAlbumView(r)}>
                    {r.album}
                    </Link>
                </div>
                <div className="card-body">
                    <div className='row'>
                        <div className='col'>
                            <Thumbnail album={r} />
                        </div>
                        <div className='col'>
                            <h4 className="card-title">
                                <Link to={linkPerformerList(r)}>
                                    {r.performer}
                                </Link>
                            </h4>
                            <p className="card-text">
                                Genre: <Link to={linkGenreList(r)}>{r.genre}</Link>
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
                    </div>

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


class Albums extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    loadData = data =>  albumViewList(data).then(res => res.data || []).then(rows => this.setState({rows}));

    componentDidMount() {
        const data = decodeURL(this.props.location.search);
        this.loadData(data);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if(JSON.stringify(prevProps.location) === JSON.stringify(this.props.location)){
            return;
        }
        this.loadData(decodeURL(this.props.location.search));
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