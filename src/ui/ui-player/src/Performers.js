import React from 'react';
import {distinctListGet} from "./api";
import { Link } from "react-router-dom";
import {linkAlbumList} from "./routes";
import { groupOnFirstChar } from "./helpers";

const PerformerLetter = ({l}) => (
    <div className='text-warning'>
        <h4>{l}</h4>
    </div>
)

const PerformerLine = ({performer}) => (
    <div>
        <Link to={linkAlbumList({performer})}>
            {performer}
        </Link>
    </div>
)

class Performers extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            items:[],
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
            .then(rows => {
                let items = [];
                rows.forEach(x => items.push(x.performer));
                this.setState({items: groupOnFirstChar(items)})
            });
    }

    render() {
        const { items } = this.state;
        return(
            <div>
                {items.map(item => (
                    <div key={item.letter} >
                        <PerformerLetter l={item.letter} />
                        {item.items.map(p => <PerformerLine key={p} performer={p} />)}
                    </div>
                ))}
            </div>
        )
    }
}

export default Performers;