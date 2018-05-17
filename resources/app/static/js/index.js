
let index = {
  init: function() {
      // Init
    //   asticode.loader.init();
    //   asticode.modaler.init();
    //   asticode.notifier.init();

      // Wait for astilectron to be ready
      document.addEventListener('astilectron-ready', function() {
          // Listen
          index.listen();

          // Explore default path
          index.explore();
      })
  },
  explore: function(path) {
      // Create message
      let message = {"name": "explore"};
      if (typeof path !== "undefined") {
          message.payload = path
      }

      // Send message
    //   asticode.loader.show();
      astilectron.sendMessage(message, function(message) {

        //   asticode.loader.hide();

          // Check error
          if (message.name === "error") {
            //   asticode.notifier.error(message.payload);
              return
          }
      })
  },
  listen: function() {
    astilectron.onMessage(function(message) {
      switch (message.name) {
      case "about":
        // index.about(message.payload);
        return {payload: "payload"};
        break;
      case "check.out.menu":
        // asticode.notifier.info(message.payload);
        break;
      }
    });
  }
};

