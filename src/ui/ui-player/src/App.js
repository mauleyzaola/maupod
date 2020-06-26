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
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Nav from "./Nav";
import { linkMediaList } from "./routes";

import { w3cwebsocket as W3CWebSocket } from "websocket";
const client = new W3CWebSocket(`ws://localhost:8080`);

class App extends React.Component{
    constructor(props) {
        super(props);
    }

    onSearch = e => {
        window.location.href = linkMediaList({query: this.query});
    }

    // stupid simple connection which is working
    componentDidMount() {
        client.onopen = () => {
            console.log('websocket connected')
        }
        client.onmessage = (message) => {
            const { data } = message;
            console.log(data);
        };
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
