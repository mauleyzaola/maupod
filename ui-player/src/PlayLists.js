import React, { Component } from "react";
import { Link } from "react-router-dom";
import API from "./api";

const ImageCard = ({r}) => {

    if(!r.image_location) return null;
    return (
        <img className="card-img-top p-0 m-0 w-40"
        src={r.image_location}
        alt="cover"/>
    )
}
const ImageCardPanel = ({r}) => {

    if(!r.image_location) return null;
    return (
        <img className="card-img" style={{width:"99px", height:"99px"}}
        src={r.image_location}
             alt="cover"/>
    )
}
const ImagePanel = ({tracks}) => {

    return (
        <>
        <div className="row">
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
            <div className="col pl-0 ml-0">
                {(tracks.count >= 3)? <ImageCardPanel r={tracks[3]}/>: ""}
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
                    <ImageCard r={playList.tracks[0]}/>:
                    <ImagePanel tracks={playList.tracks}/>}
                <div className="card-img-overlay text-white">
                    <h5 className="card-title">
                        <Link data-tip data-for="fullNameAlbum" to={playList.id}
                              title="Play">{playList.name}
                        </Link>
                    </h5>
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
            playlists:[],
        }
    }

    loadData () {
        const playlists = this.loadMockedData()
        this.setState({playlists})
        // this should be resolved in the server, not here
        // TODO: create a gh issue with the parameters you need, so we return in one request, all the information

        // let aPlayList = [];
        //    API.playLists()
        //       .then(data => {
        //         data.map(items => {
        //             API.playListItemsGet({id:items.id})
        //                .then(medias => {
        //                     if (medias !== null) {
        //                     const distImage = [...new Set(medias.map(item => item.media.image_location))];
        //                           items.tracks = distImage.filter(imagen => imagen !== undefined);
        //                           aPlayList.push(items);
        //                 }
        //                })
        //         })
        //         this.setState({playLists: aPlayList});
        //     })
            
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
        // const playLists = this.loadMockedData();
        // this.setState({playLists});
    }
    

    render(){
        const {playlists} = this.state;
        
        return(
            <>
            <div className="card-columns col-6">
                 {playlists.map(p => <CardBody key={p.id} playList={ p } />)}
            </div>
            </>
        )
    }
    
}

export default PlayLists;