package wjsimulatorsvc

type SimulateReq struct {
	Measure int    `json:"measure"`
	Jugs    []*Jug `json:"jugs"`
}

type SimulateResp struct {
	Error   bool        `json:"error"`
	Payload interface{} `json:"payload"`
}