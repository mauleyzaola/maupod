import React from 'react';
import PropTypes from 'prop-types';
import {decodeURL, directoryRead} from "../api";
import {Link} from 'react-router-dom';
import { FaFolder} from "react-icons/fa/index";

const FileList = ({files}) => (
    <table className='table table-border small'>
    <thead>
    <tr>
        <th>Name</th>
        <th>Size</th>
    </tr>
    </thead>
        <tbody>
        {files.map(x => <FileRow key={x.id} file={x} />)}
        </tbody>
    </table>
)

const FileRow = ({file}) => (
    <tr>
        <td title={file.location}>
            {file.is_dir
                ? <span> <FaFolder/>  <Link to={`/file-browser?root=${file.location}`}>{` ${file.name}`}</Link></span>
                : <span>{file.name}</span>
            }
        </td>
        <td>{!file.is_dir ? file.size : null}</td>
    </tr>
)

class FileBrowser extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            files: [],
        }
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if(JSON.stringify(prevProps.location) === JSON.stringify(this.props.location)){
            return;
        }
        this.loadData(decodeURL(this.props.location.search))
            .then(() => {});
    }

    componentDidMount() {
        this.loadData(decodeURL(this.props.location.search))
            .then(() => {});
    }

    loadData = async data => {
        const files = await directoryRead(data);
        this.setState({files});
    }

    render() {
        const { files } = this.state;
        return (
            <div>
                <FileList files={files} />
            </div>
        )
    }
}

FileBrowser.propTypes = {
    root: PropTypes.string,
}

export default FileBrowser;