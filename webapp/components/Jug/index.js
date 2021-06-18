import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

import WineGlassEmpty from '../WineGlassEmpty/';
import WineGlassFull from '../WineGlassFull/';
import WineGlassPartialFull from '../WineGlassPartialFull/';

const Jug = ({name, state, stateLabel, current}) => {
	let icon = (<WineGlassEmpty className="h-20 w-20" />);
	switch(state) {
	case "partial_full":
		icon = (<WineGlassPartialFull className="h-20 w-20" />);
		break;
	case "full":
		icon = (<WineGlassFull className="h-20 w-20" />);
		break;
	}
	
	return (
	<div className="flex flex-col justify-center items-center space-y-1">
		{icon}
		<div className="font-bold">{name}</div>
		<div className="font-light">{stateLabel}</div>
		<div className="font-light">{current}</div>
	</div>
	);
};

Jug.propTypes = {
	name: PropTypes.string,
	state: PropTypes.string,
	stateLabel: PropTypes.string,
	current: PropTypes.number,
};

Jug.defaultProps = {
	state: 'empty',
	current: 0,
};

export default Jug;