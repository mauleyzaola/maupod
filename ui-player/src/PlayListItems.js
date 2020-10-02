
import React from "react";
import propTypes from 'prop-types';
import API from "./api";

const ItemsHeader = ({id}) => {

   
    return (
        <div>
            <h1>PLAY LIST</h1>
            <div>
                <h2>{ id }</h2>
            </div>
             {/* <h2>{
                 API.playListsGet({id:id}).then(
                     response => {
                         console.log(response.data.name);
                     }
                 )
                 }</h2>  */}
        </div>
    )
} 

const MediaListHeader = () => {
    return (
        <thead>
            <tr>
                <th>#</th>
                <th>Album</th>
                <th>Title</th>
                <th>Performer</th>
                <th>Genre</th>
            </tr>
        </thead>
    )
}

const MediaListRow = ({item}) => {
    return(
        <tr>
            <td>{item.id}</td>
            <td>{item.media.album}</td>
            <td>{item.media.title}</td>
            <td>{item.media.performer}</td>
            <td>{item.media.genre}</td>
        </tr>
    )
}

class  PlayListItems extends React.Component{

    state = {
        playListId:'',
        items:[],
    }

    componentDidMount(){
        const {id} = API.decodeURL(window.location.search);
        
         this.loadData({id});
    }

    loadData = ({id}) => {
        
        API.playListItemsGet({id})
            .then( response => {
                if (response.data !== undefined)
                {
                    // let tname = ''
                    const playListId = id;
                    const items = response.data;   
                    this.setState({playListId, items})
                    }
                }
            )
    }

    render()
    {
    const {playListId, items} = this.state;
        console.log(`renderX ${playListId}`);
     return(
        <div>
            <div><ItemsHeader id={playListId}/></div>
            
             <table className='table table-bordered table-hover table-striped'>
            
              <MediaListHeader />
              <tbody> 
            {items.map(item => <MediaListRow item={item} />)}
            </tbody> 
            </table>
        </div>
     )
    }
}

ItemsHeader.propTypes = {
    id: propTypes.string,
}
MediaListRow.propTypes = {
     item: propTypes.object.isRequired,
}

export default PlayListItems;

