import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

const Input = ({label, primaryColor, helperText, error, onChange, ...restProps}) => {

	let secondaryColor = 'gray';
	if(error) {
		primaryColor = 'red';
		secondaryColor = 'red';
	}
	let fieldExt = '';
	if(error) {
		fieldExt = 'ring-1 ring-red-600 ';
	}	

	return (
	<div className="flex flex-col justify-center items-start space-y-2">
		<label className="font-bold text-gray-500 block mb-1 flex-none">
			{label}
		</label>
		<input className={`block outline-none w-full py-2 ${fieldExt} bg-white border border-${secondaryColor}-300 rounded-md shadow-md focus:ring-2 focus:ring-${primaryColor}-600 focus:border-transparent transition-all duration-500 ease-in-out pr-2 pl-3`}
			{...restProps}
			onChange={onChange}
		/>
		{ helperText && 
			<div className="flex flex-row w-full justify-start">
				<p className={`mt-1 font-extralight text-sm text-${secondaryColor}-500`}>
					{helperText}
				</p>				
			</div>
		}
	</div>		
	);
};

Input.propTypes = {
	label: PropTypes.string,
	primaryColor: PropTypes.string,
	helperText: PropTypes.string,
	error: PropTypes.bool,
	onChange: PropTypes.func,	
};

Input.defaultProps = {	
	primaryColor: 'blue',
	error: false,
};

export default Input;