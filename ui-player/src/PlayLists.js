import React, { Component } from "react";
import { Link } from "react-router-dom";
import API from "./api";

const ImageCard = ({r}) => {

    if(!r.image_location) return null;
    return (
        <img className="card-img-top p-0 m-0 w-40 "
        src={r.image_location}
        alt="cover"/>
    )
}
const ImageCardPanel = ({r}) => {

    if(!r.image_location) return null;
    return (
        <img className="card-img no-rounded" style={{width:"99px", height:"99px"}}
        src={r.image_location}
             alt="cover"/>
    )
}
const ImagePanel = ({tracks}) => {

    return (
        <>
        <div className="row ">
            <div className="col pr-0 mr-0">
                {<ImageCardPanel r={tracks[0]}/>}
            </div>
            <div className="col pl-0 ml-0">
                {<ImageCardPanel r={tracks[1]}/>}
            </div>
        </div>
        <div className="row">
            <div className="col pr-0 mr-0">
                {<ImageCardPanel r={tracks[2]}/>}
            </div>    
            <div className="col pl-0 ml-0 ">
                {(tracks.count >= 3)? <ImageCardPanel r={tracks[3]}/>: ""}
            </div>
        </div>       
        </>                        

    )
}
const CardBody = ({playList}) => {

    
    return(
            <div  className="card border-dark bg-dark p-0 pb-0 mx-2 no-rounded" 
            style={{width:"200px",height:"250px"}} >
                
                <div className="card-body p-0 m-0">
                        {(playList.tracks.length < 3)? 
                        <ImageCard r={playList.tracks[0]}/>:
                        <ImagePanel tracks={playList.tracks}/>}
                </div>
                <footer className="text-center font-weight-bold bg-transparent">
                    <div className="row pt-2">
                        <div className='col font-italic text-nowrap small'>
                            <Link data-tip  to={playList.id} title="Play">{playList.name}
                            </Link>
                        </div>
                        <div className='col font-italic text-nowrap small'>
                        {playList.tracks.length} tracks
                        </div>
                    </div>

                    <div className="row pl-4" >
                        <div className="ec-stars-wrapper">
                            <a href="#" data-value="1" title="Votar con 1 estrellas">&#9733;</a>
                            <a href="#" data-value="2" title="Votar con 2 estrellas">&#9733;</a>
                            <a href="#" data-value="3" title="Votar con 3 estrellas">&#9733;</a>
                            <a href="#" data-value="4" title="Votar con 4 estrellas">&#9733;</a>
                            <a href="#" data-value="5" title="Votar con 5 estrellas">&#9733;</a>
                        </div>
                    </div>
                    <div className="row pl-3" >
                        <Link data-tip  className="text-delete font-weight-bold p-0 m-0 small" to={playList.id} title="Delete playlist">Delete playlist
                            </Link>
                    </div>
                </footer>
            </div>
    )
}
class PlayLists extends Component{

    constructor(props) {
        super(props);
        this.state = {
            playlists:[],
        }
    }

    loadData () {
        const playlists = this.loadMockedData()
        this.setState({playlists})
            
    }

    loadMockedData() {
        const sadeLoveDeluxeCoverURL = 'https://img.discogs.com/uRbsWHELT-KMowHGWp4WJPFdGM8=/fit-in/300x300/filters:strip_icc():format(jpeg):mode_rgb():quality(40)/discogs-images/R-226306-1260746940.jpeg.jpg'
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
                        image_location : sadeLoveDeluxeCoverURL,

                     },
                     {
                        id:"2",
                        track : "You Sent Me Flying / Cherry",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                         image_location : sadeLoveDeluxeCoverURL,
                     },
                     {
                        id:"3", 
                        track : "Know You Now",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                         image_location : sadeLoveDeluxeCoverURL,
                     },
                     {
                        id:"4", 
                        track : "Fuck Me Pumps",
                        performer : "Amy Winehouse",
                        album  : "Frank",
                         image_location : sadeLoveDeluxeCoverURL,
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
                    image_location : sadeLoveDeluxeCoverURL,
                },
                // {
                //    id:"2", 
                //    track : "You Sent Me Flying / Cherry",
                //    performer : "Amy Winehouse",
                //    album  : "Frank",
                //     image_location : sadeLoveDeluxeCoverURL,
                // },
                // {
                //    id:"3", 
                //    track : "Know You Now",
                //    performer : "Amy Winehouse",
                //    album  : "Frank",
                //     image_location : sadeLoveDeluxeCoverURL,
                // },
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
                    image_location : sadeLoveDeluxeCoverURL,
                },
                {
                   id:"2", 
                   track : "You Sent Me Flying / Cherry",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                    image_location : sadeLoveDeluxeCoverURL,
                },
                {
                   id:"3", 
                   track : "Know You Now",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                    image_location : sadeLoveDeluxeCoverURL,
                },
                {
                   id:"4", 
                   track : "Fuck Me Pumps",
                   performer : "Amy Winehouse",
                   album  : "Frank",
                   image_location : "",
                }                                          
         ]              
            }
          ]
    }

    componentDidMount(){
        this.loadData()
    }
    

    render(){
        const {playlists} = this.state;
        
        return(
            <>
            <h1 className="entry-title text-center">The best playlists</h1>
            <div className="card-columns col-6">
                 {playlists.map(p => <CardBody key={p.id} playList={ p } />)}
            </div>
            </>
        )
    }
    
}

export default PlayLists;