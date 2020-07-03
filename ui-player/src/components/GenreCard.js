import {Link} from "react-router-dom";
import {linkAlbumList} from "../routes";
import {msToString} from "../helpers";
import React from "react";
import PropTypes from 'prop-types';


const Thumbnail = ({sha}) => <img className='artwork-xs col-3' alt='cover' src={`http://localhost:9000/${sha}`} />;

const pickTop = (items, top) => items.length < top ? items : items.slice(0,top);

const GenreCard = ({row}) => {
    return (
        <div className='album-card col-3'>
            <div className="card text-white bg-primary">
                <div className="card-header">
                    <Link to={linkAlbumList({genre: row.genre})}>
                        {row.genre}
                    </Link>
                </div>
                <div className="card-body">
                    <div className="row">
                        <div className="col-4 small">
                            <p className="card-text">
                                {`Performers: ${row.performer_count}`}
                            </p>
                            <p className="card-text">
                                {`Albums: ${row.album_count}`}
                            </p>
                            <p className="card-text">
                                {`Track Count: ${row.total}`}
                            </p>
                        </div>
                        <div className="col-8">
                            <div className="row">
                                {pickTop(row.artworks, 4).map(art => <Thumbnail key={art} sha={art} />)}
                            </div>
                        </div>
                    </div>
                </div>
                <div className="card-footer">
                    <small className="text-muted">
                        {row.duration ? `Duration: ${msToString(row.duration)}` : null}
                    </small>
                </div>
            </div>
        </div>
    )
}

GenreCard.propTypes = {
    row: PropTypes.object.isRequired,
}

export default GenreCard;