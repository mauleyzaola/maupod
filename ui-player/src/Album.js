import React from 'react';
import AlbumHeader from "./components/AlbumHeader";
import API from "./api";
import {msToString } from "./helpers";
import {TrackListControls} from "./components/Player";


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
    constructor(props) {
        super(props);
        this.state = {
            album:null,
            rows: [],
            genre: '',
            isCompilation: false,
        }
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

    loadData = search => {
        let album = null;
        API.albumViewList(search)
            .then(response => {
                const data = response.data || [];
                if(data.length !== 1) return;
                album  = data[0];
            })
            .then(() => API.mediaList({ sort:'track_position', direction: 'asc', ...search}))
            .then(res => res.data || [])
            .then(rows => {
                const isCompilation = this.isCompilation(rows);
                this.setState({rows, album, isCompilation})
            })
    }

    componentDidMount() {
        const uri = API.decodeURL(window.location.search);
        this.loadData(uri);
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        const { rows } = this.state;
        if(prevState.rows.length !== rows.length && rows.length !== 0){
            const genre = rows[0].genre;
            this.setState({genre});
        }
    }

    render() {
        const { album, rows, isCompilation } = this.state;
        return (
            <div>
                <AlbumHeader album={album} />
                <table className='table table-bordered table-hover table-striped'>
                    <TrackListHeader isCompilation={isCompilation}/>
                    <tbody>
                    {rows.map(row => <TrackListRow key={row.id} row={row} isCompilation={isCompilation} />)}
                    </tbody>
                </table>
            </div>
        );
    }
}

export default Album;