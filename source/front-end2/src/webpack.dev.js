const { merge } = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common, {
    mode: 'development',
    devServer: {
        contentBase: __dirname + "/public",
        // publicPath: "/js",
        port: 3001
    },
    watch: true
})
