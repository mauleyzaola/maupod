import React from 'react';
import {distinctListGet} from "./api";
import uuid from "uuid4";


const AlbumHeader = () => (
    <thead>
    <tr>
        <td>Name</td>
    </tr>
    </thead>
)

function AlbumLine({row}){
    return (
        <tr>
            <td>{row.album}</td>
        </tr>
    )
}

class Albums extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        distinctListGet({
            field:'album',
            filter:{
                direction: 'asc',
                sort: 'album',
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
                    <AlbumHeader />
                    <tbody>
                    {rows.map(row => <AlbumLine key={uuid()} row={row}  />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Albums;