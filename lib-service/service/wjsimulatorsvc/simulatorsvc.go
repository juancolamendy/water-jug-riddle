package wjsimulatorsvc

import (
	"log"

	"github.com/juancolamendy/water-jug-riddle/lib-service/utils/errorcodes"
	"github.com/juancolamendy/water-jug-riddle/lib-service/utils/mathutils"
)

type SimulatorSvcOpts struct {
	Verbose bool
}

func DefaultOpts() *SimulatorSvcOpts {
	return &SimulatorSvcOpts{Verbose: true}
}

type SimulatorSvc struct {
	inChan chan *SimulateReq
	outChan chan *SimulateResp

	isProcessing bool

	verbose bool
}

func NewSimulatorSvc(opts *SimulatorSvcOpts) *SimulatorSvc {
	svc := &SimulatorSvc{
		inChan: make(chan *SimulateReq),
		outChan: make(chan *SimulateResp),
		
		isProcessing: false,

		verbose: opts.Verbose,
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

func (s *SimulatorSvc) buildJugsStatus(lastStep bool, bigJug *Jug, smallJug *Jug) *JugsStatus {
	jugMap := make(map[string]Jug)

	jugMap[bigJug.Name] = *bigJug
	jugMap[smallJug.Name] = *smallJug

	return &JugsStatus {
		LastStep: lastStep,
		JugMap: jugMap,
	}
}

func (s *SimulatorSvc) doSimulation(req *SimulateReq) {
	// Init
	var bigJug *Jug
	var smallJug *Jug
	
	if req.Jugs[0].Capacity > req.Jugs[1].Capacity {
		bigJug = req.Jugs[0]
		smallJug = req.Jugs[1]
	} else {
		bigJug = req.Jugs[1]
		smallJug = req.Jugs[0]
	}
	bigJug.empty()
	smallJug.empty()
	if s.verbose {
		log.Println("Start simulation")
		log.Printf("Measure: %d\n", req.Measure)
		smallJug.dump()
		bigJug.dump()
	}
	
	// Logic
	s.outChan <- &SimulateResp {
		Error: false,
		Payload: s.buildJugsStatus(false, bigJug, smallJug),
	}
	
	bigJug.fill()
	s.outChan <- &SimulateResp {
		Error: false,
		Payload: s.buildJugsStatus(false, bigJug, smallJug),
	}	
	if s.verbose {
		log.Println("initial fill bigJug")
		bigJug.dump()
	}	

	for {		
		bigJug.transferTo(smallJug)
		status := s.buildJugsStatus(bigJug.Current == req.Measure, bigJug, smallJug)
		s.outChan <- &SimulateResp {
			Error: false,
			Payload: status,
		}
		if s.verbose {
			log.Println("transferTo from bigJug to smallJug")
			bigJug.dump()
			smallJug.dump()
		}
		if status.LastStep {
			break
		}

		if bigJug.Current == 0 {			
			bigJug.fill()
			status := s.buildJugsStatus(bigJug.Current == req.Measure, bigJug, smallJug)
			s.outChan <- &SimulateResp {
				Error: false,
				Payload: status,
			}
			if s.verbose {
				log.Println("fill bigJug")
				bigJug.dump()
			}
			if status.LastStep {
				break
			}			
		}

		if smallJug.Current == smallJug.Capacity {			
			smallJug.empty()
			status := s.buildJugsStatus(smallJug.Current == req.Measure, bigJug, smallJug)
			s.outChan <- &SimulateResp {
				Error: false,
				Payload: status,
			}
			if s.verbose {
				log.Println("empty smallJug")
				smallJug.dump()
			}
			if status.LastStep {
				break
			}			
		}
	}
}

func (s *SimulatorSvc) eventLoop() {
	for {
		select {
		case req := <- s.inChan:
			// Log req
			log.Printf("--- received request: %+v", req)

			// Validate req
			if ok, msg := s.validateReq(req); !ok {
				log.Printf("--- sending error response: %+v", msg)
				s.outChan <- &SimulateResp {
					Error: true,
					Payload: msg,
				}
				continue
			}

			// Start processing
			s.isProcessing = true
			log.Printf("--- start processing: %+v", req)

			// Do simulation
			s.doSimulation(req)

			// End processing
			s.isProcessing = false
			log.Printf("--- end processing: %+v", req)
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