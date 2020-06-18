import {Link} from "react-router-dom";
import {linkAlbumList} from "../routes";
import {msToString} from "../helpers";
import React from "react";

const GenreCard = ({r}) => {
    return (
        <div className='album-card col-3'>
            <div className="card text-white bg-primary">
                <div className="card-header">
                    <Link to={linkAlbumList({genre: r.genre})}>
                        {r.genre}
                    </Link>
                </div>
                <div className="card-body">
                    <p className="card-text">
                        {`Performers: ${r.performer_count}`}
                    </p>
                    <p className="card-text">
                        {`Albums: ${r.album_count}`}
                    </p>
                    <p className="card-text">
                        {`Track Count: ${r.total}`}
                    </p>

                </div>
                <div className="card-footer">
                    <small className="text-muted">
                        {r.duration ? `Duration: ${msToString(r.duration)}` : null}
                    </small>
                </div>
            </div>
        </div>
    )
}

export default GenreCard;