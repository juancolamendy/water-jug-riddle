package wjsimulatorsvc

// Example:
// {"measure": 4, "jugs": [{"capacity": 5, "name": "jug01"}, {"capacity": 3, "name": "jug02"}] }
type SimulateReq struct {
	Measure int    `json:"measure"`
	Jugs    []*Jug `json:"jugs"`
}

type SimulateResp struct {
	Error   bool        `json:"error"`
	Payload interface{} `json:"payload"`
}