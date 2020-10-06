import React, { Component } from "react";
import API from "./api";

const CardImagen = ({image_location}) => {

    if(!image_location) return null;
    return (
        <img className="card-img-top p-2 w-30"
            src={`${process.env.REACT_APP_MAUPOD_ARTWORK}/${image_location}`} 
            alt="  "/>
    )
}
const ListGroup = ({track}) => {
    return (
        <a href="#" className="list-group-item list-group-item-action flex-column align-items-start bg-dark active px-2 py-0 mx-0">
            <p className="my-0">{ track.track }</p>
            <small className="py-0"> {track.performer} / {track.album}</small>
        </a> 
    )
}

const CardBody = ({playList}) => {
    return (
            <div  className="card border-secondary bg-dark p-2 pb-0 mx-2 no-rounded" style={{width:"220px"}} >
                {<CardImagen image_location={ playList.imagen_location }/>}
                <div className="card-body p-1 m-0">
                    <h3 className="card-title text-center">{ playList.name }</h3>
                </div>

                <small>Include</small>
                <div className="list-group p-0 mx-0">
                    { playList.tracks.map(track => <ListGroup key={track.id} track={track} />) }
                 </div>
                <footer>Play</footer>
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

    loadData = data =>  API.playLists().then(res => res.data || []).then(playLists => this.setState({playLists}));

    loadMockedData() {
        return [
            {
              id: "9e196558-1bed-4b77-8dc8-867c87985fed",
              name: "Only Rock",
              imagen_location : "4dd32f59-fbd0-48a3-b3c1-574dd3130353.png",
              tracks: [
                     {
                        id:"1",
                        track : "Intro / Stronger Than Me",
                        performer : "Amy Winehouse",
                        album  : "Frank",

                     },
                     {
                        id:"2",
                        track : "You Sent Me Flying / Cherry",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                     },
                     {
                        id:"3", 
                        track : "Know You Now",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                     },
                     {
                        id:"4", 
                        track : "Fuck Me Pumps",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                     }                                          
              ]
            },
            {
              id: "7316c65a-4837-4244-87ea-c93c165e0163",
              name: "Only Rock1",
              imagen_location : "32612fe4-5b4a-42ff-824c-0ccb83d63eec.png",
              tracks: [
                {
                   id:"1",  
                   track : "Intro / Stronger Than Me",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"2", 
                   track : "You Sent Me Flying / Cherry",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"3", 
                   track : "Know You Now",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"4", 
                   track : "Fuck Me Pumps",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                }                                          
         ]              
            },
            {
              id: "2d2a5d55-5154-4679-aeb5-bcd772486617",
              name: "Only Rock2",
              imagen_location: "3ad3b977-97d6-49bc-a88d-313fa496acba.png",
              tracks: [
                {
                   id:"1", 
                   track : "Intro / Stronger Than Me",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"2", 
                   track : "You Sent Me Flying / Cherry",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"3", 
                   track : "Know You Now",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                },
                {
                   id:"4", 
                   track : "Fuck Me Pumps",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                }                                          
         ]              
            }
          ]
    }

    componentDidMount(){
        //this.LoadData()
        const playLists = this.loadMockedData();
        this.setState({playLists});
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