package wjsimulatorsvc

import (
	"log"

	"github.com/juancolamendy/water-jug-riddle/apisvr/utils/errorcodes"
	"github.com/juancolamendy/water-jug-riddle/apisvr/utils/mathutils"
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

func (s *SimulatorSvc) buildJugsStatus(bigJug Jug, smallJug Jug) *JugsStatus {
	jugMap := make(map[string]Jug)

	jugMap[bigJug.Name] = bigJug
	jugMap[smallJug.Name] = smallJug

	return &JugsStatus {
		JugMap: jugMap,
	}
}

func (s *SimulatorSvc) pour(measure int, sourceJug Jug, targetJug Jug) []*JugsStatus {
	// Init
	result := make([]*JugsStatus, 0)

	// Logic
	sourceJug.empty()
	targetJug.empty()
	if s.verbose {
		log.Println("start simulation")
		log.Printf("measure: %d\n", measure)
		targetJug.dump()
		sourceJug.dump()
	}
	result = append(result, s.buildJugsStatus(sourceJug, targetJug))
	
	sourceJug.fill()	
	if s.verbose {
		log.Println("initial fill sourceJug")
		sourceJug.dump()
	}
	result = append(result, s.buildJugsStatus(sourceJug, targetJug))

	for {		
		sourceJug.transferTo(&targetJug)
		result = append(result, s.buildJugsStatus(sourceJug, targetJug))
		if s.verbose {
			log.Println("transferTo from sourceJug to targetJug")
			sourceJug.dump()
			targetJug.dump()
		}
		if sourceJug.Current == measure || targetJug.Current == measure {
			break
		}

		if sourceJug.Current == 0 {			
			sourceJug.fill()
			result = append(result, s.buildJugsStatus(sourceJug, targetJug))
			if s.verbose {
				log.Println("fill sourceJug")
				sourceJug.dump()
			}
			if sourceJug.Current == measure || targetJug.Current == measure {
				break
			}			
		}

		if targetJug.Current == targetJug.Capacity {			
			targetJug.empty()
			result = append(result, s.buildJugsStatus(sourceJug, targetJug))
			if s.verbose {
				log.Println("empty targetJug")
				targetJug.dump()
			}
			if sourceJug.Current == measure || targetJug.Current == measure {
				break
			}			
		}
	}

	return result
}

func (s *SimulatorSvc) doSimulation(req *SimulateReq) []*JugsStatus {
	// Run simulations async
	sim1Ch := make(chan []*JugsStatus)
	go func(ch chan []*JugsStatus) {
		result := s.pour(req.Measure, *req.Jugs[0], *req.Jugs[1])		
		ch <- result
	}(sim1Ch)

	sim2Ch := make(chan []*JugsStatus)
	go func(ch chan []*JugsStatus) {
		result := s.pour(req.Measure, *req.Jugs[1], *req.Jugs[0])		
		ch <- result
	}(sim2Ch)

	// Wait and get result
	result1 := <- sim1Ch
	result2 := <- sim2Ch

	log.Printf("end simulation 1 - steps: %d", len(result1))
	log.Printf("end simulation 2 - steps: %d", len(result2))
	if len(result1) > len(result2) {
		return result2
	}
	return result1
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
			result := s.doSimulation(req)
			s.outChan <- &SimulateResp {
				Error: false,
				Payload: result,
			}

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