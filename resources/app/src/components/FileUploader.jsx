import React, { Component } from 'react';
import FileSelector from './FileSelector';

class FileUploader extends Component {
  constructor(props) {
    super(props);

    this.state = {
      path: '',
      keyPath: '',
    }

    this.uploadFile = this.uploadFile.bind(this);
    this.onFilePathChanged = this.onFilePathChanged.bind(this);
  }

  uploadFile() {
    if (!this.state.path) {
      return;
    }
    const message = {
      name: 'uploadFile',
      payload: {
        Path: this.state.path,
        KeyPath: this.state.keyPath
      }
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
        <div className="section">
          <div className="heading">Вивантажити файл</div>
          <div>Файл для вивантаження: </div>
          <FileSelector onFilePathChanged={(p) => this.onFilePathChanged('path', p)}/>

          <div className="button" onClick={this.uploadFile}>Вивантажити файл</div>
        </div>
      </div>
    );
  }
}

export default FileUploader;
