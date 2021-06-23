import React, { useState, useRef, useEffect } from 'react';

const useDelayCall = (force, fnCallback, delay) => {

	const isInit = useRef(true);

	useEffect(() => {
		if(isInit.current) {
			isInit.current = false;
			return;
		}

		const handler = setTimeout(() => {
			fnCallback();
		}, delay);

		return () => {
			clearTimeout(handler);
		};

	}, [force]);

};

export default useDelayCall;