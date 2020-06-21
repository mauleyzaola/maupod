import React from 'react';
import 'bootswatch/dist/slate/bootstrap.min.css';
import './App.css';
import Albums from "./Albums";
import Album from "./Album";
import AudioScan from "./AudioScan";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import {BrowserRouter as Router, Link, Route, Switch} from 'react-router-dom';
import Nav from "./Nav";
import { linkMediaList } from "./routes";

class App extends React.Component{
    onSearch = e => {
        const uri = linkMediaList({query: this.query});
        window.location.href = uri;
    }

    onSearchChange = e => this.query = e.target.value;

    render() {
        return (
            <div>
                <Router>
                    <Nav onSearch={this.onSearch} onChange={this.onSearchChange} />
                    <Switch>
                        <Route exact path='/' component={Dashboard} />
                        <Route exact path='/audio-scan' component={AudioScan} />
                        <Route exact path='/albums' component={Albums} />
                        <Route path='/album' component={Album} />
                        <Route exact path='/performers' component={Performers} />
                        <Route exact path='/genres' component={Genres} />
                        <Route path='/media' component={MediaList} />
                    </Switch>
                </Router>
            </div>
        );
    }
}

export default App;
