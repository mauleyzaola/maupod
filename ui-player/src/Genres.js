import React from 'react';
import API from "./api";
import GenreCard from "./components/GenreCard";


class Genres extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        let rows = [];
        API.genreList({
            direction: 'asc',
            sort: 'genre',
        })
            .then(res => res.data || [])
            .then(data => rows = data)
            .then(() => API.genreArtworkList())
            .then(response => {
                const artworks = response.data;
                for(let i = 0; i < rows.length; i++){
                    let row = rows[i];
                    row.artworks = artworks[row.genre] || [];
                }
                this.setState({rows});
            })
    }

    render() {
        const { rows } = this.state;
        return(
            <div className='card-deck'>
                {rows.map(row => <GenreCard key={row.genre} row={row} />)}
            </div>
        )
    }
}

export default Genres;