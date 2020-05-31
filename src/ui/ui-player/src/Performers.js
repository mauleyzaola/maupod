import React from 'react';
import {distinctListGet} from "./api";
import uuid from 'uuid4';
import { Link } from "react-router-dom";
import {linkPerformerList} from "./routes";

const PerformerHeader = () => (
    <thead>
        <tr>
            <td>Name</td>
        </tr>
    </thead>
)

function PerformerLine({row}){
    const { performer } = row;
    return (
        <tr>
            <td>
                <Link to={linkPerformerList(row)}>
                    {performer}
                </Link>
            </td>
        </tr>
    )
}

class Performers extends React.Component{
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
                    <PerformerHeader />
                    <tbody>
                        {rows.map(row => <PerformerLine key={uuid()} row={row}  />)}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Performers;