import React from 'react';
import { connect } from 'react-redux';
import 'bootswatch/dist/slate/bootstrap.min.css';
import './App.css';
import Albums from "./Albums";
import Album from "./Album";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import {Route, Switch} from 'react-router-dom';
import Nav from "./Nav";
import { linkMediaList } from "./routes";
import FileBrowser from "./components/FileBrowser";
import TrackControl from "./components/TrackControl";
import Queue from "./components/Queue";
import Setup from "./components/Setup";

class App extends React.Component{
    onSubmit = e => {
        window.location.href = linkMediaList({query: this.query});
    }
    onSearchChange = e => this.query = e.target.value;

    render() {
        return (
            <div className='container-fluid'>
                <div className='navigation-bar'>
                    <Nav
                        onSubmit={this.onSubmit}
                        onChange={this.onSearchChange}
                    />
                </div>
                <div className='top-always'>
                    <TrackControl />
                </div>

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
                        <Route path='/setup' component={Setup} />
                    </Switch>
                </div>
            </div>
        );
    }
}

export default connect((state) => ({ }))(App);