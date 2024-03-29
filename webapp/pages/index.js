import React, { useState, useMemo, useEffect } from "react";

import { Alert, Input, Button, Jug } from '../components/';

import { useDelayCall, useWebsocket } from '../hooks/';

import constant from '../utils/constant';
import locales from '../utils/locales';
import { buildBaseWsUrl } from '../utils/webutils';

const baseWsUrl = buildBaseWsUrl();
const WS_URL = `${baseWsUrl}${constant.API_WS_INDEX}`;

const buildRequest = ({measure, amount1, amount2}) => JSON.stringify({
	"measure": parseInt(measure), 
	"jugs": [{"capacity": parseInt(amount1), "name": "jug1"}, {"capacity": parseInt(amount2), "name": "jug2"}],
	"ts": Date.now(),
});

const Index = () => {

	// State
	const [forceOpen, setForceOpen] = useState(false);
	const [forceShowResult, setForceShowResult] = useState(false);

	const [processing, setProcessing] = useState(false);
	const [result, setResult] = useState(null);
	const [error, setError] = useState(null);
	const [amount1, setAmount1] = useState(5);
	const [amount2, setAmount2] = useState(3);
	const [measure, setMeasure] = useState(4);

	// Computed values
	const amount1Invalid = useMemo(() => isNaN(amount1) || amount1 <= 0, [amount1]);
	const amount2Invalid = useMemo(() => isNaN(amount2) || amount2 <= 0, [amount2]);
	const measureInvalid = useMemo(() => isNaN(measure) || measure <= 0, [measure]);
	const buttonValid = useMemo(() => !amount1Invalid && !amount2Invalid && !measureInvalid, [amount1Invalid, amount2Invalid, measureInvalid])

	// Effects
	const { setWsInput, wsOutput, wsConnected } = useWebsocket(WS_URL, forceOpen);

	useEffect(() => {
		if(!wsConnected) {
			console.log('reconnect');
			setForceOpen(!forceOpen);
		}
	}, [wsConnected]);

	useEffect(() => {
		const data = JSON.parse(wsOutput);
		console.log('received:', data);
		if(data) {
			if(data.error) {
				setError(locales.errors[data.payload]);
				setForceShowResult(!forceShowResult);
			} else {
				setResult(data.payload);
				setForceShowResult(!forceShowResult);				
			}
		}
	}, [wsOutput]);

	useDelayCall(forceShowResult, () => {
		setProcessing(false);
	}, constant.DELAY);		

	// Handlers
	const handleSubmit = (event) => {
		event.preventDefault();

		const data = buildRequest({measure, amount1, amount2});
		console.log('sending request:', data);
		setError(null);
		setResult(null);
		setProcessing(true);
		setWsInput(data);		
	};

	const handleCloseError = () => {
		setError(null);
	};

	return (
	<div className="flex flex-col h-full items-center p-10 bg-gray-200">

		<div className="flex flex-col flex-grow h-full lg:w-3/12 mb-5">
			<h1 className="font-extralight text-center text-3xl">
				{locales.app_title}
			</h1>

			{ error && !processing &&
			<div className="mt-2">
				<Alert caption={locales.label_error}
					description={error}
					onClick={handleCloseError}
				/>				
			</div>
			}

			<div className="mt-5 p-8 bg-white rounded-lg shadow-lg border-blue-500 border-t-8 space-y-4">
				<form className="flex flex-col space-y-3"
					onSubmit={handleSubmit} 
					noValidate>
					<Input label={locales.label_jug1_capacity}
						value={amount1}
						helperText={locales.hint_input_num}
						error={amount1Invalid}
						type="number"
						onChange={evt => setAmount1(evt.target.value)}
					/>	
					<Input label={locales.label_jug2_capacity}
						value={amount2}
						helperText={locales.hint_input_num}
						error={amount2Invalid}
						type="number"
						onChange={evt => setAmount2(evt.target.value)}
					/>
					<Input label={locales.label_to_measure}
						value={measure}
						helperText={locales.hint_input_num}
						error={measureInvalid}
						type="number"
						onChange={evt => setMeasure(evt.target.value)}
					/>
					<Button valid={buttonValid}
						label={locales.label_simulate}
						processing={processing}
						processingLabel={locales.label_simulating}
						onClick={handleSubmit}				
					/>
				</form>
			</div>
		</div>

		{result && !processing &&
		result.map((x,i) => <div className="flex flex-row justify-between m-3 border-b-2 border-gray-600 pb-2" key={i}>
			<Jug name="jug1" state={x.jugMap['jug1'].state} stateLabel={locales.state[x.jugMap['jug1'].state]} current={x.jugMap['jug1'].current} />
			<Jug name="jug2" state={x.jugMap['jug2'].state} stateLabel={locales.state[x.jugMap['jug2'].state]} current={x.jugMap['jug2'].current} />
		</div>)
		}

	</div>
	);	
};

export default Index;