import React from 'react';
import 'bootswatch/dist/slate/bootstrap.min.css';
import './App.css';
import Albums from "./Albums";
import AudioScan from "./AudioScan";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Nav from "./Nav";

class App extends React.Component{
    onSearch = e => {
        e.preventDefault();
        alert('not implemented yet!');
    }

    render() {
        return (
            <div>
                <Router>
                    <Nav onSearch={this.onSearch} />
                    <Switch>
                        <Route exact path='/' component={Dashboard} />
                        <Route exact path='/audio-scan' component={AudioScan} />
                        <Route exact path='/albums' component={Albums} />
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
