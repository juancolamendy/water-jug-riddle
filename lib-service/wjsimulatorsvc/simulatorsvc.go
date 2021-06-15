package wjsimulatorsvc

import (
	"log"

	"github.com/juancolamendy/water-jug-riddle/lib-service/utils/errorcodes"
	"github.com/juancolamendy/water-jug-riddle/lib-service/utils/mathutils"
)

type SimulatorSvc struct {
	inChan chan *SimulateReq
	outChan chan *SimulateResp

	isProcessing bool
}

func NewSimulatorSvc() *SimulatorSvc {
	svc := &SimulatorSvc{
		inChan: make(chan *SimulateReq),
		outChan: make(chan *SimulateResp),
		
		isProcessing: false,
	}
	go svc.eventLoop()

	return svc
}

func (s *SimulatorSvc) validateReq(req *SimulateReq) (bool, string) {
	// Validate jugs
	if len(req.Jugs) <= 0 || len(req.Jugs) > 2 {
		return false, errorcodes.ERROR_ONLY_TWO_JUGS_ALLOWED
	}
	if len(req.Jugs[0].Name) == 0 || len(req.Jugs[1].Name) == 0 {
		return false, errorcodes.ERROR_MUST_PROVIDE_JUG_NAMES
	}
	if req.Jugs[0].Capacity <= 0 || req.Jugs[1].Capacity <= 0 {
		return false, errorcodes.ERROR_JUG_CAPACITY_MUST_GREATER_0
	}

	// Validate measure
	if req.Measure <= 0 {
		return false, errorcodes.ERROR_MEASURE_MUST_GREATER_0
	}
	maxCap := mathutils.Max(req.Jugs[0].Capacity, req.Jugs[1].Capacity) 
	if req.Measure >= maxCap {
		return false, errorcodes.ERROR_MEASURE_MUST_LESS_MAX_CAPACITY
	}

	// Validate if there is a solution
	gcdnum := mathutils.Gcd(req.Jugs[0].Capacity, req.Jugs[1].Capacity)
	if !mathutils.IsMultiple(req.Measure, gcdnum) {
		return false, errorcodes.ERROR_NO_SOLUTION
	}

	// Return valid
	return true, ""
}

func (s *SimulatorSvc) eventLoop() {
	for {
		select {
		case req := <- s.inChan:
			// Log req
			log.Printf("--- received request: %+v", req)

			// Validate req
			if ok, msg := s.validateReq(req); !ok {
				s.outChan <- &SimulateResp {
					Error: false,
					Payload: msg,
				}
			}

			// Start processing
			s.isProcessing = true

			// Init processing structure

			// End processing
			s.isProcessing = false
		}
	}
}

func (s *SimulatorSvc) Simulate(req *SimulateReq) {
	if !s.isProcessing {
		s.inChan <- req
	}	
}

func (s *SimulatorSvc) GetOutChan() chan *SimulateResp {
	return s.outChan
}