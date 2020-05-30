import React from 'react';
import { distinctListGet } from "./api";
import uuid from 'uuid4';

const ArtistHeader = () => (
    <tr>
        <td>Name</td>
    </tr>
)

function ArtistLine({row}){
    return (
        <tr>
            <td>{row.performer}</td>
        </tr>
    )
}

class Artists extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            rows:[],
        }
    }

    componentDidMount() {
        distinctListGet({
            field:'performer',
            filter:{
                direction: 'asc',
                sort: 'performer',
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
                    <thead>
                        <ArtistHeader />
                    </thead>
                    <tbody>
                        {rows.map(row => <ArtistLine key={uuid()} row={row}  />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Artists;