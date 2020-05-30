import React from 'react';
import './App.css';
import Albums from "./Albums";
import Performers from './Performers'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import MediaList from "./MediaList";
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';

function App() {
  return (
    <div>
        <Router>
            <div>
                <Link to='/'>Home</Link>
                <Link to='/albums'>Albums</Link>
                <Link to='/performers'>Performers</Link>
                <Link to='/genres'>Genres</Link>
            </div>
            <Switch>
                <Route exact path='/' component={Dashboard} />
                <Route exact path='/albums' component={Albums} />
                <Route exact path='/performers' component={Performers} />
                <Route exact path='/genres' component={Genres} />
                <Route path='/media' component={MediaList} />
            </Switch>
        </Router>
    </div>
  );
}

export default App;
