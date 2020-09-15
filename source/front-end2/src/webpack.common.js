// https://stackoverflow.com/questions/36039146/webpack-dev-server-compiles-files-but-does-not-refresh-or-make-compiled-javascri

module.exports = {
    mode: 'development',
    // context: __dirname + "/game",
    entry: "./game/index.ts",
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/
            }
        ]
    },
    devtool: 'source-map',
    output: {
        path: __dirname + "/dist",
        // publicPath: "/dist",
        filename: "bundle.js"
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    }
}
