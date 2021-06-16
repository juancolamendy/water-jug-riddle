package wjsimulatorsvc

import (
	"testing"
	"log"
)

func printOutJugs(jugMap map[string]Jug) {
	for _, v := range jugMap {
		v.dump()
	}
}

func TestSimulator(t *testing.T) {
	opts := DefaultOpts()
	opts.Verbose = false
	svc := NewSimulatorSvc(opts)
	
	jug01 := &Jug {
		Name: "jug01",
		Capacity: 11,
	}
	jug02 := &Jug {
		Name: "jug02",
		Capacity: 9,
	}
	req := &SimulateReq{
		Measure: 4,
		Jugs: []*Jug {jug01, jug02},
	}
	svc.Simulate(req)

	for resp := range svc.GetOutChan() {
		log.Printf("--- received response: %+v", resp)
		if jugMap, ok := resp.Payload.(map[string]Jug); ok {
			printOutJugs(jugMap)
		}		
	}
}