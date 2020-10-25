import React from 'react';
import AlbumHeader from "./components/AlbumHeader";
import API from "./api";
import {msToString } from "./helpers";
import {TrackListControls} from "./components/Player";
// import { useParams } from 'react-router-dom'

const CoverLine = ({c, onClick}) => (
    <div className="cover-item">
        <div className="card">
            <img src={c.cover_image} alt="..." width="200" onClick={() => onClick(c)} />
        </div>
    </div>
)

const TrackListHeader = ({isCompilation}) => {
    if(isCompilation) {
        return (
            <thead>
            <tr>
                <th>#</th>
                <th>Track</th>
                <th>Performer</th>
                <th>Duration</th>
                <th>Format</th>
            </tr>
            </thead>
        )
    }
    return (
        <thead>
        <tr>
            <th>#</th>
            <th>Track</th>
            <th>Duration</th>
            <th>Format</th>
        </tr>
        </thead>
    )
}

const TrackListRow = ({isCompilation, row}) => {
    if(isCompilation){
        return (
            <tr>
                <td>{row.track_position}</td>
                <td>
                    <div className='row'>
                        <div className='col-4'>
                            <TrackListControls media={row} />
                        </div>
                        <div className='col-8'>
                            {row.track}
                        </div>
                    </div>
                </td>
                <td>{row.performer}</td>
                <td>{msToString(row.duration)}</td>
                <td>{row.format}</td>
            </tr>
        )
    }
    return (
        <tr>
            <td>{row.track_position}</td>
            <td>
                <div className='row'>
                    <div className='col-4'>
                        <TrackListControls media={row} />
                    </div>
                    <div className='col-8'>
                        {row.track}
                    </div>
                </div>
            </td>
            <td>{msToString(row.duration)}</td>
            <td>{row.format}</td>
        </tr>
    )
}


class Album extends React.Component{
    state = {
        album:null,
        rows: [],
        genre: '',
        isCompilation: false,
        covers: [],
    }

    isCompilation = (rows) => {
        let performer = '';
        for(let i = 0; i < rows.length; i++){
            const row = rows[i];
            if(i === 0){
                performer = row.performer;
                continue;
            }
            if(performer !== row.performer){
                return true;
            }
        }
        return false;
    }

    loadData = id => {
        const data = {album_identifier: id}
        let album = null;
        API.albumViewList(data)
            .then(response => {
                const data = response.data || [];
                if(data.length !== 1) return;
                album  = data[0];
            })
            .then(() => API.mediaList({ sort:'track_position', direction: 'asc', ...data}))
            .then(res => res.data || [])
            .then(rows => {
                const isCompilation = this.isCompilation(rows);
                this.setState({rows, album, isCompilation})
            })
    }

    componentDidMount() {
        const { id } = this.props.match.params
        this.loadData(id);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        const { id } = this.props.match.params
        const { album } = this.state
        if(album.album_identifier && album.album_identifier !== id){
            this.loadData(id)
        }
    }

    onCoverFormSubmit = e => {
        e.preventDefault();
        const { rows } = this.state;
        if(rows.length === 0) return;
        let { performer: artist, recorded_date: year, album: title } = rows[0];
        const type = 'master';
        const data = { artist, year, title, type };
        API.providerMetadataCovers({params: data})
            .then(response => {
                let { covers } = this.state;
                covers = response.data;
                this.setState({covers});
                if(covers.length === 0){
                    // TODO: make this better, like a notification alert in the ui
                    alert(`provider returned no cover images`);
                }
            })
            .catch(error => {
                if(error.response && error.response.data){
                    // TODO: make this better, like a notification alert in the ui
                    alert(JSON.stringify(error.response.data))
                }
            })
    }

    onCoverClick = c => {
        const { album  } = this.state;
        let { album_identifier } = album;
        API.providerMetadataCoverPut({
            params: {
                album_identifier,
            },
            data:{
                uri: c.cover_image,
                force: true, // overwrite current artwork if exists
            }
        })
            .then(response => {
                const { album } = this.state;
                album.image_location = `${album.album_identifier}.png`;
                this.setState({album});
            } )
            .catch(error => {
                if(error.response && error.response.data){
                    // TODO: make this better, like a notification alert in the ui
                    alert(JSON.stringify(error.response.data))
                }
            })
    }


    render() {
        const { album, covers, rows, isCompilation } = this.state;
        return (
            <div>
                <AlbumHeader album={album} onClick={this.onCoverFormSubmit} />
                <div className="row">
                    <div className="col">
                        {covers.length !== 0 &&
                        <div className="card-group">
                            {covers.map(c => <CoverLine key={c.uuid} c={c} onClick={this.onCoverClick} />)}
                        </div>
                        }
                    </div>
                </div>
                <div className="row">
                    <div className="col">
                        <table className='table table-bordered table-hover table-striped'>
                            <TrackListHeader isCompilation={isCompilation}/>
                            <tbody>
                            {rows.map(row => <TrackListRow key={row.id} row={row} isCompilation={isCompilation} />)}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        );
    }
}

export default Album;