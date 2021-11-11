const path = require('path');
const mode = process.env.MODE ?? 'production';

module.exports = {
  entry: './src/index.ts',
  mode,
  target: 'web',
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        use: 'ts-loader',
        exclude: /node_modules/,
      },
    ],
  },
  resolve: {
    extensions: ['.ts', '.js'],
  },
  output: {
    filename: 'trie.js',
    path: path.resolve(__dirname, 'build'),
    library: "tsTrie"
  }
};