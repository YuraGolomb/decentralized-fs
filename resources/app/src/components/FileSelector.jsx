import React from 'react';

class FileSelector extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      filepath: ''
    }

    this.onFilePathChange = this.onFilePathChange.bind(this);
    this.onFileSelect = this.onFileSelect.bind(this);
  }

  onFilePathChange(e) {
    this.setState({ filepath: e.target.value });
    this.props.onFilePathChanged(e.target.value)
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
      <div className="file-uploader-wrapper">
        <input
          className="file-path"
          type="text"
          onChange={this.onFilePathChange}
          value={this.state.filepath}
        />
        <div
          className="file-button"
          type="button"
          onClick={this.onFileSelect}
        >Select File</div>
      </div>
    );
  }
}

export default FileSelector;
