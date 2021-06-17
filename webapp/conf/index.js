// Environment variables
const NODE_ENV = process.env.NODE_ENV;
const LOG_LEVEL = process.env.LOG_LEVEL || 'debug';
const PORT = parseInt(process.env.PORT, 10) || 3000;

const API_SCHEME = process.env.API_SCHEME || 'http';
const API_HOST   = process.env.API_HOST   || 'localhost';
const API_PORT   = process.env.API_PORT   || '3001';

const dumpConfig = () => {
	console.log('--- Dumping config');

	console.log(`--- NODE_ENV: ${NODE_ENV}`);
	console.log(`--- LOG_LEVEL: ${LOG_LEVEL}`);
	console.log(`--- PORT: ${PORT}`);

	console.log(`--- API_SCHEME: ${API_SCHEME}`);
	console.log(`--- API_HOST: ${API_HOST}`);
	console.log(`--- API_PORT: ${API_PORT}`);
};

// Export
exports = module.exports = {
	NODE_ENV,
	LOG_LEVEL,
	PORT,

	API_SCHEME,
	API_HOST,
	API_PORT,

	dumpConfig,	
};