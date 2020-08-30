import React from 'react';

class Setup extends React.Component{
    render() {
        return (
            <div>
                <h3>Events</h3>
                <form action={`${process.env.REACT_APP_MAUPOD_API}/events`} method="get">
                    <button className='btn btn-info' type="button">
                        Export Events
                    </button>
                </form>
                <form
                    action={`${process.env.REACT_APP_MAUPOD_API}/events`}
                    method="post"
                    enctype="multipart/form-data"
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