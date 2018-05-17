const path = require('path');
module.exports = {
  entry: { main: './resources/app/src/app.js' },
  output: {
    path: path.resolve('./resources/app/static/js'),
    filename: 'app.js'
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader']
      }
    ]
  },
  resolve: {
    extensions: ['*', '.js', '.jsx']
  },
};
