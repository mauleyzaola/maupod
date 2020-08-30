import React from 'react';

class Setup extends React.Component{
    render() {
        return (
            <div>
                <h3>Events</h3>
                <a
                    href={`${process.env.REACT_APP_MAUPOD_API}/events`}
                    className="btn btn-info"
                    rel="noopener noreferrer"
                    target="_blank">
                    Export Events
                </a>
                <form
                    action={`${process.env.REACT_APP_MAUPOD_API}/events`}
                    method="post"
                    encType="multipart/form-data"
                    target="_blank">
                    <input type="file" name="file" />
                    <button className='btn btn-warning' type="submit">
                        Import Events
                    </button>
                </form>
            </div>
        )
    }
}

export default Setup;