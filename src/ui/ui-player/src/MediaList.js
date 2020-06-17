import React from 'react';
import { decodeURL, mediaList } from "./api";
import { TrackListHeader, TrackListRow } from "./components/TrackList";

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
                <table className='table table-bordered table-hover'>
                    <TrackListHeader />
                    <tbody>
                    {rows.map(row => <TrackListRow key={row.id} row={row} />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default MediaList;