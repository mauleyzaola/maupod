import React from 'react';
import {distinctListGet, objectToQueryString} from "./api";
import uuid from "uuid4";
import {Link} from "react-router-dom";


const GenreHeader = () => (
    <thead>
    <tr>
        <td>Name</td>
    </tr>
    </thead>
)

function GenreLine({row}){
    const { genre } = row;
    return (
        <tr>
            <td>
                <Link to={`/media?${objectToQueryString({genre})}`}>
                    {genre}
                </Link>
            </td>
        </tr>
    )
}

class Genres extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        distinctListGet({
            field:'genre',
            filter:{
                direction: 'asc',
                sort: 'genre',
            },
        })
            .then(res => res.data || [])
            .then(rows => this.setState({rows}));
    }

    render() {
        const { rows } = this.state;
        return(
            <div>
                <table>
                    <GenreHeader />
                    <tbody>
                    {rows.map(row => <GenreLine key={uuid()} row={row}  />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Genres;