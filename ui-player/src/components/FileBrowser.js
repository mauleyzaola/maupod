import React from 'react';
import PropTypes from 'prop-types';
import {decodeURL, directoryRead} from "../api";
import {Link} from 'react-router-dom';
import { FaFolder} from "react-icons/fa/index";

const FileList = ({files, onClick}) => (
    <table className='table table-border small'>
    <thead>
    <tr>
        <th>Name</th>
        <th>Size</th>
    </tr>
    </thead>
        <tbody>
        {files.map(x => <FileRow key={x.id} file={x} onClick={onClick} />)}
        </tbody>
    </table>
)

const FileRow = ({file, onClick}) => {

    const css = file.selected ? 'text-warning' : '';

    return (
        <tr>
            <td className={css} title={file.location} onClick={() => onClick(file)}>
                {file.is_dir
                    ? <span> <FaFolder/>  <Link to={`/file-browser?root=${file.location}`}>{` ${file.name}`}</Link></span>
                    : <span>{file.name}</span>
                }
            </td>
            <td>{!file.is_dir ? file.size : null}</td>
        </tr>
    )
}

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

    onClick = file => {
        const { files } = this.state;
        const item = files.find(x => x.id === file.id);
        if(!item){
            return;
        }
        item.selected = !item.selected;
        this.setState({files});
    }

    render() {
        const { files } = this.state;
        return (
            <div>
                <FileList files={files} onClick={this.onClick} />
            </div>
        )
    }
}

FileBrowser.propTypes = {
    root: PropTypes.string,
}

export default FileBrowser;