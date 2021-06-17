export const buildBaseWsUrl = () => {
	if(typeof window !== 'undefined') {
		let protocol = 'ws';
		const host = window.location.host;
		if (window.location.protocol === 'https:') protocol = 'wss';

		return `${protocol}://${host}`;
	}
	return '';
};