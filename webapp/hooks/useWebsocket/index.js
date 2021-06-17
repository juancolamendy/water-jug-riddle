import React, { useEffect, useRef, useState } from 'react';

const useWebsocket = (url, forceOpen) => {

	// Refs
	const ws = useRef(null);

	// State
	const [wsInput, setWsInput] = useState(null);
	const [wsOutput, setWsOutput] = useState(null);
	const [wsConnected, setWsConnected] = useState(false);

	// Effects
	useEffect(() => {
		ws.current = new WebSocket(url);
		ws.current.onopen = () => {
			console.log("websocket opened");
			setWsConnected(true);
		};
		ws.current.onclose = () => {
			console.log("websocket closed");
			setWsConnected(false);
		};
		ws.current.onmessage = evt => {
			setWsOutput(evt.data);
		};

		return () => {
			ws.current.close();
		};
	}, [url, forceOpen]);

	useEffect(() => {
		if(!ws.current || !wsConnected || ws.current.readyState != 1) return;

		ws.current.send(wsInput);
	}, [wsInput]);

	// Returns
	return { setWsInput, wsOutput, wsConnected };
};

export default useWebsocket;