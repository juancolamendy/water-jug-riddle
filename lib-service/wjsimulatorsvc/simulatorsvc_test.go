package wjsimulatorsvc

import (
	"testing"
	"log"
)

func TestSimulator(t *testing.T) {
	svc := NewSimulatorSvc()
	
	jug01 := &Jug {
		Name: "jug01",
		Capacity: 5,
	}
	jug02 := &Jug {
		Name: "jug02",
		Capacity: 3,
	}
	req := &SimulateReq{
		Measure: 4,
		Jugs: []*Jug {jug01, jug02},
	}
	svc.Simulate(req)

	for resp := range svc.GetOutChan() {
		log.Printf("--- received response: %+v", resp)
	}
}