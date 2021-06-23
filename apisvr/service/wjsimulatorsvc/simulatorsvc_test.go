package wjsimulatorsvc

import (
	"testing"
	"log"
)

func printOutJugs(status []*JugsStatus) {
	for _, s := range status {
		log.Print("--- status")
		for _, v := range s.JugMap {
			v.dump()
		}		
	}
}

func TestSimulator(t *testing.T) {
	opts := DefaultOpts()
	opts.Verbose = false
	svc := NewSimulatorSvc(opts)
	
	jug01 := &Jug {
		Name: "jug01",
		Capacity: 2,
	}
	jug02 := &Jug {
		Name: "jug02",
		Capacity: 14,
	}
	req := &SimulateReq{
		Measure: 4,
		Jugs: []*Jug {jug01, jug02},
	}
	svc.Simulate(req)

	for resp := range svc.GetOutChan() {
		log.Printf("--- received response: %+v", resp)
		if status, ok := resp.Payload.([]*JugsStatus); ok {
			printOutJugs(status)
		}		
	}
}