import React from 'react';
import API from "./api";
import {linkAlbumView, linkGenreList, linkPerformerList} from "./routes";
import { Link } from "react-router-dom";
import {msToString} from "./helpers";

const Thumbnail = ({album}) => {
    if(!album.image_location){
        //return null;

        return (
          <img alt='cover' className='artwork-small' src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/unknown.png`} />
        )
    }
    return (
        <img alt='cover' className='artwork-small' src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/${album.image_location}`} />
    )
}


const AlbumCard = ({r}) => {
    return (

        <div className='album-card col-2 ml-1 card-deck'>
            <div className="card border-secondary bg-dark p-0 mx-2 w-5 no-rounded">
                <div className='card-img-top'>      
                            <Thumbnail album={r} />
                </div>       
                <div class="card-body p-1 text bg-dark h-50">
                    <div className='row'>
                        <div className='col mx-1 font-italic text-nowrap small'>
                        {r.recorded_date ? `${r.recorded_date}` : "unknown"}
                        </div>
                        <div className='col font-italic text-nowrap small'>
                        {r.track_name_total ? `Tracks: ${r.track_name_total}` : null}
                        </div>
                    </div>               
                    <div className='mx-1 text-truncate small'>
                            <Link data-tip data-for="fullNameAlbum" to={linkAlbumView(r)} title={r.album} >
                            {r.album}</Link>
                    </div>
                    <div className='mx-1 text-muted small'>
                        <Link to={linkPerformerList(r)} title={r.performer}>
                        {r.performer}
                        </Link>     
                    </div>

                </div> 
                <div class="card-footer">
                    <div className="row">
                        <div className="col">
                            <small className="text-muted h-5">
                                {r.duration ? msToString(r.duration) : null}
                            </small>
                        </div>
                        <div className="col">
                            <small className="text-muted h-5">
                                <Link to={linkGenreList(r)} title={r.genre}>
                                    {r.genre}
                                </Link>
                            </small>
                        </div>
                    </div>
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

    loadData = data =>  API.albumViewList(data).then(res => res.data || []).then(rows => this.setState({rows}));

    componentDidMount() {
        const data = API.decodeURL(this.props.location.search);
        if(!data.limit){
            data.limit = 100;
        }
        this.loadData(data);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if(JSON.stringify(prevProps.location) === JSON.stringify(this.props.location)){
            return;
        }
        this.loadData(API.decodeURL(this.props.location.search));
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