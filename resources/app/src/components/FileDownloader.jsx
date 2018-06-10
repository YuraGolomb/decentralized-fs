import React, { Component } from 'react';
import FileSelector from './FileSelector';

class FileUploader extends Component {
  constructor(props) {
    super(props);

    this.state = {
      path: '',
      keyPath: ''
    }

    this.uploadFile = this.uploadFile.bind(this);
    this.onFilePathChanged = this.onFilePathChanged.bind(this);
  }

  uploadFile() {
    if (!this.state.path) {
      return;
    }
    console.log(this.state.path);
    const message = {
      name: 'downloadFile',
      payload: {
        Path: this.state.path,
        KeyPath: this.state.keyPath
      }
      // payload: JSON.stringify({
      //   Path: this.state.path,
      //   KeyPath: this.state.keyPath
      // }),
    };
    astilectron.sendMessage(message, function(message) {
      console.log(message);
    })
  }

  onFilePathChanged(type, path) {
    this.setState({ [type]: path });
  }

  render() {
    return (
      <div>
        <div>File:</div>
        <FileSelector onFilePathChanged={(p) => this.onFilePathChanged('path', p)}/>
        <div>Key:</div>
        <FileSelector onFilePathChanged={(p) => this.onFilePathChanged('keyPath', p)}/>

        <div onClick={this.uploadFile}>DownloadFile</div>
      </div>
    );
  }
}

export default FileUploader;
