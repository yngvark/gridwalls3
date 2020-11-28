// https://stackoverflow.com/questions/36039146/webpack-dev-server-compiles-files-but-does-not-refresh-or-make-compiled-javascri
const path = require('path')

var phaserModule = path.join(__dirname, '/node_modules/phaser/')
var phaser = path.join(phaserModule, 'dist/phaser.js')

module.exports = {
    // context: __dirname + "/game",
    // entry: ['./game/index.ts'],
    entry: {
        main: "./game/index.ts",
        vendor: ['phaser']
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                include: path.join(__dirname, '../src'),
                use: 'ts-loader',
            }
        ]
    },
    devtool: 'source-map',
    output: {
        path: __dirname + "/dist",
        publicPath: "/dist",
        filename: "[name].bundle.js",
        chunkFilename: '[name].chunk.js'
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js'],
        alias: {
            'phaser': phaser,
        }
    },
    optimization: {
        splitChunks: {
            cacheGroups: {
                commons: {
                    test: /[\\/]node_modules[\\/]/,
                    name: 'vendors',
                    chunks: 'all',
                    filename: '[name].bundle.js'
                }
            }
        }
    },
}
