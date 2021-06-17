const path = require('path');
const express = require('express');
const next = require('next');
const { createProxyMiddleware } = require('http-proxy-middleware');

const conf = require('./conf');

conf.dumpConfig();

const port = conf.PORT;
const dev = conf.NODE_ENV !== 'production';

const app = next({ dev });
const handle = app.getRequestHandler();

const apiServerUrl = `${conf.API_SCHEME}://${conf.API_HOST}:${conf.API_PORT}`;
console.log(`API_SERVER_URL: ${apiServerUrl}`);

app.prepare().then(() => {
	const server = express();
 
	server.use(
		'/api/v1',
		createProxyMiddleware({
			target: apiServerUrl,
			logLevel: conf.LOG_LEVEL,
			changeOrigin: true,
			ws: true,
		})
	);

	server.get('*', (req, res) => {
		return handle(req, res);
	});

	console.log(`Trying to listen on port: [${port}]`);
	server.listen(port, err => {
		if (err) throw err;
		console.log(`Ready to serve on port: [${port}]`)
	});
});