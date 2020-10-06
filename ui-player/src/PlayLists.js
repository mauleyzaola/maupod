import React, { Component } from "react";
import propTypes from 'prop-types';
import API from "./api";

class PlayLists extends Component{

    constructor(props) {
        super(props);
        this.state = {
            items:[],
        }
    }

    loadData = data =>  API.playLists().then(res => res.data || []).then(items => this.setState({items}));

    componentDidMount(){
        //this.LoadData()
        document.title = "PlayLists";
        
    }
    

    render(){

        const {items} = this.state;

        return(
            <>
           <div class="card-columns col-6">
    
                {/* <div className="card border-secondary bg-dark p-2 pb-0 mx-2 no-rounded" style={{width:"250px", height:"600px"}} >
              
                    <div className="card-header p-0 m-0">
                    <h5 className="card-title text-center">Only Rock</h5>
                    </div>
                    <img className="card-img-top p-2 w-30" src="http://192.168.1.77:7401/4dd32f59-fbd0-48a3-b3c1-574dd3130353.png" alt="Card image cap"></img>
                    <p className="card-text">Some quick example .</p>
                    <small>Include</small>
                    <div className="card-body pl-0 m-0">
                     <div className="list-group p-0 mx-0">
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">you Know I'm not good.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Me & Mr jones.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Just Friends.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Read.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>                                                            
                  </div>
                </div>
                   <div className="card-footer bg-transparent border-success"> Rating</div> 
                </div>  */}

          
                <div style={{width:"220px"}} className="card border-secondary bg-dark p-2 pb-0 mx-2 no-rounded">
                    <img className="card-img-top p-2 w-30" src="http://192.168.1.77:7401/4dd32f59-fbd0-48a3-b3c1-574dd3130353.png" alt="Card image cap"></img>

                    <div className="card-body p-1 m-0">
                        <h3 className="card-title text-center">Only Rock</h3>
                        {/* <p className="card-text">Some quick example .</p> */}
                    </div>

                <small>Include</small>
                <div className="list-group p-0 mx-0">
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start bg-dark active px-2 py-0 mx-0">
                    <p className="my-0">you Know I'm not good.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a> 
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start bg-dark active px-2 py-0 mx-0">
                        <p className="my-0">Me & Mr jones.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start bg-dark active px-2 py-0 mx-0">
                        <p className="my-0">Just Friends.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start bg-dark active px-2 py-0 mx-0">
                        <p className="my-0">Read.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>                                                            
                </div>
                <footer>Play</footer>
                </div> 

                <div style={{width:"220px"}} className="card border-secondary bg-dark p-2 pb-0 mx-2 no-rounded">
                    <img className="card-img-top p-2 w-30" src="http://192.168.1.77:7401/4dd32f59-fbd0-48a3-b3c1-574dd3130353.png" alt="Card image cap"></img>

                    <div className="card-body p-1 m-0">
                        <h3 className="card-title text-center">Urban Rock</h3>
                        <p className="card-text">Some quick example .</p>
                    </div>

                <small>Include</small>
                <div className="list-group p-0 mx-0">
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">you Know I'm not good.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Me & Mr jones.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Just Friends.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>    
                    <a href="#" className="list-group-item list-group-item-action flex-column align-items-start active px-2 py-0 mx-0">
                        <p className="my-0">Read.</p>
                        <small className="py-0">Auto / Disco</small>
                    </a>                                                            
                </div>
                <footer>Play</footer>
                </div> 
         
            </div>
            
   
            </>
        )
    }
    
}

export default PlayLists;