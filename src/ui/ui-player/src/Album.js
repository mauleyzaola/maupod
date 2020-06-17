import React from 'react';
import AlbumHeader from "./components/AlbumHeader";
import {decodeURL} from "./api";
import { mediaList } from "./api";

class Album extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            album:null,
        }
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
        const { album } = this.state;
        return (
            <div>
                <AlbumHeader props={album} />
            </div>
        );
    }
}

export default Album;