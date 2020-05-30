import React from 'react';
import './App.css';
import Albums from "./Albums";
import Artists from './Artists'
import Genres from "./Genres";
import Dashboard from "./Dashboard";
import { BrowserRouter as Router, Link, Route, Switch } from 'react-router-dom';

function App() {
  return (
    <div>
        <Router>
            <div>
                <Link to='/'>Home</Link>
                <Link to='/albums'>Albums</Link>
                <Link to='/artists'>Artists</Link>
                <Link to='/genres'>Genres</Link>
            </div>
            <Switch>
                <Route exact path='/' component={Dashboard} />
                <Route exact path='/albums' component={Albums} />
                <Route exact path='/artists' component={Artists} />
                <Route exact path='/genres' component={Genres} />
            </Switch>
        </Router>
    </div>
  );
}

export default App;