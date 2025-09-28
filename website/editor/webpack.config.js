const HtmlWebpackPlugin = require('html-webpack-plugin');
const CompressionPlugin = require("compression-webpack-plugin");
const path = require('path');

module.exports = {
	entry: './src/index.tsx',
	mode: "production",
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'main.[contenthash].js',
	},
	devServer: {
		static: {
			directory: path.resolve(__dirname, 'dist'),
		},
		port: 3000,
		open: true,
		hot: true,
		compress: true,
		historyApiFallback: true,
	},
	module: {
		rules: [
			{
				test: /\.tsx?$/,
				use: 'ts-loader',
				exclude: /node_modules/,
			},
			{
				test: /\.?js$/,
				exclude: /node_modules/,
				use: {
					loader: "babel-loader",
					options: {
						presets: [
							"@babel/preset-env",
							["@babel/preset-react", { "runtime": "automatic" }]
						]
					}
				}
			},
			{
				test: /\.(css)$/,
				use: ['style-loader', 'css-loader', 'postcss-loader'],
			},
		],
	},
	resolve: {
		extensions: ['.tsx', '.ts', '.js'],
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: "./src/index.html",
		}),
		new CompressionPlugin(),
	],
};
