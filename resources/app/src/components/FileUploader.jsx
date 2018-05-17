import React, { Component } from 'react';
import FileSelector from './FileSelector';

class FileUploader extends Component {
  constructor(props) {
    super(props);

    this.state = {
      filePath: '',
    }

    this.uploadFile = this.uploadFile.bind(this);
    this.onFilePathChanged = this.onFilePathChanged.bind(this);
  }

  uploadFile() {
    if (!this.state.filePath) {
      return;
    }
    console.log(this.state.filePath);
    const message = { name: 'uploadFile'};
    message.payload = this.state.filePath;
    astilectron.sendMessage(message, function(message) {
      console.log(message);
    })
  }

  onFilePathChanged(filePath) {
    this.setState({ filePath });
  }

  render() {
    return (
      <div>
        <FileSelector onFilePathChanged={this.onFilePathChanged}/>

        <div onClick={this.uploadFile}>UploadFile</div>
      </div>
    );
  }
}

export default FileUploader;
