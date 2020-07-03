import { Link } from "react-router-dom";
import React from "react";
import PropTypes from 'prop-types';

class Nav extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            query: '',
        }
    }


    onChange = e => {
        if(this.props.onChange){
            this.props.onChange(e);
        }
        this.setState({query: e.target.value})
    }

    render() {
        const { query } = this.state;
        return (
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
                    <form className="form-inline my-2 my-lg-0" onSubmit={e => {
                        e.preventDefault();
                        this.props.onSearch(e);
                        this.setState({query:''});
                    }}>
                        <input
                            className="form-control mr-sm-2" type="text" placeholder="Search"
                            onChange={this.onChange}
                            value={query}
                        />
                        <button className="btn btn-secondary my-2 my-sm-0"
                                disabled={query.length === 0}
                                type="submit">Search</button>
                    </form>
                </div>
            </nav>
        )
    }
}

export default Nav;

Nav.propTypes = {
    onChange: PropTypes.func.isRequired,
    onSearch: PropTypes.func.isRequired,
}