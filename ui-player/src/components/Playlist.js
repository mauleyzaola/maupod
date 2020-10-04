import React from "react";
import propTypes from 'prop-types';

const ItemsHeader = ({playlist}) => (
    <div>
        <h1>{playlist.name}</h1>
    </div>
)

const MediaListHeader = () => (
    <thead>
    <tr>
        <th>Album</th>
        <th>Track</th>
        <th>Performer</th>
        <th>Genre</th>
    </tr>
    </thead>
)

const MediaListRow = ({item}) => (
    <tr>
        <td>{item.media.album}</td>
        <td>{item.media.track}</td>
        <td>{item.media.performer}</td>
        <td>{item.media.genre}</td>
    </tr>
)

class Playlist extends React.Component{
    state = {
        id:'',
        playlist:{},
        items:[],
    }

    componentDidMount(){
        const { id, onLoadData } = this.props;
        const { playlist, items} = onLoadData(id);
        this.setState({playlist, items});
    }

    render() {
    const {playlist, items} = this.state;
     return(
        <div>
            <div>
                <ItemsHeader playlist={playlist} />
            </div>
            
             <table className='table table-bordered table-hover table-striped'>
              <MediaListHeader />
              <tbody> 
                {items.map(item => <MediaListRow key={item.id} item={item} />)}
              </tbody>
            </table>
        </div>
     )
    }
}

Playlist.propTypes = {
    id: propTypes.string.isRequired,
    onLoadData: propTypes.func.isRequired,
}

export default Playlist;

