import React from 'react';
import 'bootswatch/dist/slate/bootstrap.min.css';
import './App.css';
import Albums from "./Albums";
import AudioScan from "./AudioScan";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';

class App extends React.Component{
    render() {
        return (
            <div>
                <Router>
                    <nav className="navbar navbar-expand-lg navbar-dark bg-primary">
                        <Link className="navbar-brand" to='/'>Maupod</Link>
                        <div className="collapse navbar-collapse" id="navbarColor01">
                            <ul className="navbar-nav mr-auto">
                                <li className={`nav-item`}>
                                    <Link className="nav-link" to='/albums'>Albums</Link>
                                </li>
                                <li className="nav-item">
                                    <Link className="nav-link" to='/genres'>Genres</Link>
                                </li>
                                <li className="nav-item">
                                    <Link className="nav-link" to='/performers'>Performers</Link>
                                </li>
                                <li className="nav-item">
                                    <Link className="nav-link" to='/audio-scan'>Audio Scan</Link>
                                </li>
                            </ul>
                            <form className="form-inline my-2 my-lg-0">
                                <input className="form-control mr-sm-2" type="text" placeholder="Search" />
                                <button className="btn btn-secondary my-2 my-sm-0" type="submit">Search</button>
                            </form>
                        </div>
                    </nav>
                    <div>

                    </div>
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
