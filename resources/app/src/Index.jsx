import React from 'react';

import FileUploader from './components/FileUploader.jsx';
import FileDownloader from './components/FileDownloader.jsx';

export default class Index extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};

    this.explore = this.explore.bind(this);
    this.listen = this.listen.bind(this);
  }

  componentDidMount() {

    const self = this;
    document.addEventListener('astilectron-ready', function() {
      self.listen();
      self.explore();
    })
  }

  explore(path) {
    // Create message
    let message = {"name": "init"};
    if (typeof path !== "undefined") {
        message.payload = path
    }

    // Send message
    astilectron.sendMessage(message, function(message) {
        console.log(message)

        // Check error
        if (message.name === "error") {
            return
        }
    })
  }

  listen() {
    astilectron.onMessage(function(message) {
      switch (message.name) {
      case "about":
        console.log(mesasge);
        // return {payload: "payload"};
        break;
      case "check.out.menu":
        console.log(message)
        break;
      }
    })
  }


  render() {
    return (
      <div>
        Here is React app!!
        <FileUploader />
        <FileDownloader />
      </div>
    )
  }
}
