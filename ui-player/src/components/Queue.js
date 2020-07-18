import React from 'react';
import PropTypes from 'prop-types';

class Queue extends React.Component{
    render() {
        return (
            <div>
                <button className='btn btn-sm btn-info'>
                    Play Now
                </button>
                <button className='btn btn-sm btn-secondary'>
                    Add to queue
                </button>
                <button className='btn btn-sm btn-warning'>
                    Play Next
                </button>
            </div>
        )
    }
}

Queue.propTypes = {
    isVisible: PropTypes.bool.isRequired,
}

export default Queue;