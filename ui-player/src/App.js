import React from 'react';
import 'bootswatch/dist/slate/bootstrap.min.css';
import './App.css';

import Albums from "./Albums";
import Album from "./Album";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Nav from "./Nav";
import { linkMediaList } from "./routes";
import FileBrowser from "./components/FileBrowser";
import TrackControl from "./components/TrackControl";
import Queue from "./components/Queue";

class App extends React.Component{

    onSearch = e => {
        // this is doggy shit, page gets refreshed on each submission
        window.location.href = linkMediaList({query: this.query});
    }


    onSearchChange = e => this.query = e.target.value;

    render() {
        return (
            <div className='container-fluid'>
                <Router>
                    <div className="top-half">
                        <div className='navigation-bar'>
                            <Nav
                                onSearch={this.onSearch}
                                onChange={this.onSearchChange}
                            />
                        </div>
                        <div className='top-always'>
                            <TrackControl />
                        </div>
                    </div>
                    <div className="bottom-half">
                        <div className='scroll-section'>
                            <Switch>
                                <Route exact path='/' component={Dashboard} />
                                <Route exact path='/albums' component={Albums} />
                                <Route path='/album' component={Album} />
                                <Route exact path='/performers' component={Performers} />
                                <Route exact path='/genres' component={Genres} />
                                <Route path='/media' component={MediaList} />
                                <Route path='/queue' component={Queue} />
                                <Route path='/file-browser' component={FileBrowser} />
                            </Switch>
                        </div>
                    </div>    
                </Router>
            </div>
        );
    }
}

export default App;
