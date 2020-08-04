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
                    <Nav
                        onSearch={this.onSearch}
                        onChange={this.onSearchChange}
                    />
                    <Switch>
                        <Route exact path='/' component={Dashboard} />
                        <Route exact path='/albums' component={Albums} />
                        <Route path='/album' component={Album} />
                        <Route exact path='/performers' component={Performers} />
                        <Route exact path='/genres' component={Genres} />
                        <Route path='/media' component={MediaList} />
                        <Route path='/file-browser' component={FileBrowser} />
                    </Switch>
                </Router>
                <div>
                    <TrackControl />
                </div>
            </div>
        );
    }
}

export default App;
