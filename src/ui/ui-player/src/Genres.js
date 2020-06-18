import React from 'react';
import {genreList} from "./api";
import GenreCard from "./components/GenreCard";


class Genres extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        genreList({
            direction: 'asc',
            sort: 'genre',
        })
            .then(res => res.data || [])
            .then(rows => this.setState({rows}));
    }

    render() {
        const { rows } = this.state;
        return(
            <div className='card-deck'>
                {rows.map(r => <GenreCard key={r.genre} r={r} />)}
            </div>
        )
    }
}

export default Genres;