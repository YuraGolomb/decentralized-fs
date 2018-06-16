import React from 'react';
import FileInfo from './FileInfo';

class FileSelector extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      filepath: ''
    }

    this.onFilePathChange = this.onFilePathChange.bind(this);
    this.onFileSelect = this.onFileSelect.bind(this);
    this.onFileInfo = this.onFileInfo.bind(this);
  }

  onFilePathChange(e) {
    this.setState({ filepath: e.target.value });
    this.props.onFilePathChanged(e.target.value)
  }

  onFileInfo() {
    if (!this.state.filepath) {
      return;
    }
    const message = {
      name: 'checkFile',
      payload: {
        Path: this.state.filepath,
        KeyPath: ' '
      }
    };
    const self = this;
    astilectron.sendMessage(message, function(message) {
      console.log(message);
      self.setState({ file: message.payload });
    })
    
  }

  onFileSelect() {
    const self = this;
    astilectron.showOpenDialog({ properties: ['openFile'], title: "Select File" }, function(paths) {
      console.log("chosen paths are ", paths)
      self.setState({ filepath: paths[0] });
      self.props.onFilePathChanged(paths[0])
    })
  }

  render() {
    return (
      <div>
        <div className="file-uploader-wrapper">
          <input
            className="file-path"
            type="text"
            className="text-input"
            onChange={this.onFilePathChange}
            value={this.state.filepath}
          />
          <div
            className="file-button"
            type="button"
            className="button"
            onClick={this.onFileSelect}
          >Обрати файл</div>
          <div
            className="file-button"
            type="button"
            className="button"
            onClick={this.onFileInfo}
          >Інформація</div>
        </div>
        { this.state.file && <FileInfo file={this.state.file} />}
      </div>
    );
  }
}

export default FileSelector;
