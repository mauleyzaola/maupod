import React, { Component } from "react";
import { Link } from "react-router-dom";
import API from "./api";

const CardImagen = ({image_location}) => {

    if(!image_location) return null;
    return (
        <img className="card-img-top p-0 m-0 w-40"
        src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/${image_location}`} 
        alt="  "/>
    )
}
const CardImagenPanel = ({image_location}) => {

    if(!image_location) return null;
    return (
        <img className="card-img" style={{width:"99px", height:"99px"}}
        src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/${image_location}`} 
        alt="  "/>
    )
}
const PanelImagen = ({tracks}) => {

    return (
        <>
        <div className="row">
            <div className="col pr-0 mr-0">
                {<CardImagenPanel image_location={tracks[0]}/>}
            </div>
            <div className="col pl-0 ml-0">
                {<CardImagenPanel image_location={tracks[1]}/>}
            </div>
        </div>
        <div className="row">
            <div className="col pr-0 mr-0">
                {<CardImagenPanel image_location={tracks[2]}/>}
            </div>    
            <div className="col pl-0 ml-0">
                {(tracks.count >= 3)? <CardImagenPanel image_location={tracks[3]}/>: ""}
            </div>
        </div>       
        </>                        

    )
}
const CardBody = ({playList}) => {

    return(
        <div  className="card border-dark bg-dark p-0 pb-0 mx-2 no-rounded" style={{width:"200px",height:"250px"}} >
            <div className="card-body p-0 m-0">
                    {(playList.tracks.length < 3)? 
                    <CardImagen image_location={playList.tracks[0]}/>: 
                    <PanelImagen tracks={playList.tracks}/>}
                <div className="card-img-overlay text-white">
                    <h5 className="card-title">
                        <Link data-tip data-for="fullNameAlbum" to={playList.id} title="Play">{playList.name}</Link></h5>
                    <p className="card-text">Last updated 3 mins ago</p>
                </div>
            </div>
            <footer className="text-center font-weight-bold text-info">
                <div className="row">
                    
                        <div className="col">
                        <Link data-tip  to={playList.id} title="Play">{playList.tracks.length} tracks
                        </Link> 
                    </div>
                </div>
            </footer>
     </div>

    )
}
class PlayLists extends Component{

    constructor(props) {
        super(props);
        this.state = {
            playLists:[],
        }
    }

    loadData () {
        let aPlayList = [];
           API.playLists()
              .then(data => {
                data.map(items => { 
                    API.playListItemsGet({id:items.id})
                       .then(medias => {
                            if (medias !== null) {
                            const distImage = [...new Set(medias.map(item => item.media.image_location))];
                                  items.tracks = distImage.filter(imagen => imagen !== undefined);
                                  aPlayList.push(items);
                        }
                       })
                }) 
                this.setState({playLists: aPlayList});
            })
            
    }

    loadMockedData() {
        return [
            {
              id: "9e196558-1bed-4b77-8dc8-867c87985fed",
              name: "Only Rock",
              tracks: [
                     {
                        id:"1",
                        track : "Intro / Stronger Than Me",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                        imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",

                     },
                     {
                        id:"2",
                        track : "You Sent Me Flying / Cherry",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                        imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                     },
                     {
                        id:"3", 
                        track : "Know You Now",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                        imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                     },
                     {
                        id:"4", 
                        track : "Fuck Me Pumps",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                        imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                     }                                          
              ]
            },
            {
              id: "7316c65a-4837-4244-87ea-c93c165e0163",
              name: "Musica Colombiana",
              tracks: [
                {
                   id:"1",  
                   track : "Intro / Stronger Than Me",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "32612fe4-5b4a-42ff-824c-0ccb83d63eec.png",
                },
                {
                   id:"2", 
                   track : "You Sent Me Flying / Cherry",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "32612fe4-5b4a-42ff-824c-0ccb83d63eec.png",
                },
                {
                   id:"3", 
                   track : "Know You Now",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                },
                // {
                //    id:"4", 
                //    track : "Fuck Me Pumps",
                //    performer : "Amy Winehouse",
                //    album  : "Frank",
                //    imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                // }                                          
         ]              
            },
            {
              id: "2d2a5d55-5154-4679-aeb5-bcd772486617",
              name: "Rock Urbano",
              tracks: [
                {
                   id:"1", 
                   track : "Intro / Stronger Than Me",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
                },
                {
                   id:"2", 
                   track : "You Sent Me Flying / Cherry",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location: "3ad3b977-97d6-49bc-a88d-313fa496acba.png",
                },
                {
                   id:"3", 
                   track : "Know You Now",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "32612fe4-5b4a-42ff-824c-0ccb83d63eec.png",
                },
                {
                   id:"4", 
                   track : "Fuck Me Pumps",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   imagen_location : "",
                }                                          
         ]              
            }
          ]
    }

    componentDidMount(){
        this.loadData()
        // const playLists = this.loadMockedData();
        // this.setState({playLists});
    }
    

    render(){
        const {playLists} = this.state;
        
        return(
            <>
            <div className="card-columns col-6">
                 {playLists.map(playList => <CardBody key={playList.id} playList={ playList } />)}
            </div>
            </>
        )
    }
    
}

export default PlayLists;