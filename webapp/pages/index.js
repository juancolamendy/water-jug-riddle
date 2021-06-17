import React, { useState, useMemo, useEffect } from "react";

import { Input, Button } from '../components/';

import { useWebsocket } from '../hooks/';

const Index = () => {

	// State
	const [forceOpen, setForceOpen] = useState(false);

	const [processing, setProcessing] = useState(false);
	const [amount1, setAmount1] = useState(5);
	const [amount2, setAmount2] = useState(3);
	const [measure, setMeasure] = useState(4);

	// Computed values
	const amount1Invalid = useMemo(() => amount1 <= 0, [amount1]);
	const amount2Invalid = useMemo(() => amount2 <= 0, [amount2]);
	const measureInvalid = useMemo(() => measure <= 0, [measure]);
	const buttonValid = useMemo(() => !amount1Invalid && !amount2Invalid && !measureInvalid, [amount1Invalid, amount2Invalid, measureInvalid])

	// Effects
	const { setWsInput, wsOutput, wsConnected } = useWebsocket('ws://localhost:3000/api/v1/ws', forceOpen);

	useEffect(() => {
		if(!wsConnected) {
			console.log('reconnect');
			setForceOpen(!forceOpen);
		}
	}, [wsConnected]);

	useEffect(() => {
		const data = JSON.parse(wsOutput);
		console.log('received:', data);			
	}, [wsOutput]);

	// Handlers
	const handleSubmit = (event) => {
		event.preventDefault();
		console.log('handleSubmit:');

		const data = JSON.stringify({
			"measure": parseInt(measure), 
			"jugs": [{"capacity": parseInt(amount1), "name": "jug01"}, {"capacity": parseInt(amount2), "name": "jug02"}],
			"ts": Date.now(),
		});
		setWsInput(data);		
	};

	return (
	<div className="flex flex-col h-screen items-center pt-10 bg-gray-200">
		<h1 className="font-extralight text-center text-3xl">
			Water Jug Riddle
		</h1>
		<div className="mt-5 p-8 bg-white rounded-lg shadow-lg border-blue-500 border-t-8 space-y-4">
			<form className="flex flex-col space-y-3"
				onSubmit={handleSubmit} 
				noValidate>
				<Input label="* Jug 1 Gallons"
					value={amount1}
					helperText="Amount of Gallons. Greater than 0"
					error={amount1Invalid}
					type="number"
					onChange={evt => setAmount1(evt.target.value)}
				/>	
				<Input label="* Jug 2 Gallons" 
					value={amount2}
					helperText="Amount of Gallons. Greater than 0"
					error={amount2Invalid}
					type="number"
					onChange={evt => setAmount2(evt.target.value)}
				/>
				<Input label="* Measure in Gallons"
					value={measure}
					helperText="Amount Greater than 0"
					error={measureInvalid}
					type="number"
					onChange={evt => setMeasure(evt.target.value)}
				/>
				<Button valid={buttonValid}
					label="Simulate"
					processing={processing}
					processingLabel="Simulating ..."
					onClick={handleSubmit}				
				/>
			</form>
		</div>
	</div>
	);	
};

export default Index;