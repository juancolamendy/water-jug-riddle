package wjsimulatorsvc

import (
	"log"

	"github.com/juancolamendy/water-jug-riddle/apisvr/utils/mathutils"
)

const (
	STATE_EMPTY        = "empty"
	STATE_FULL         = "full"
	STATE_PARTIAL_FULL = "partial_full"
)

type Jug struct {
	Capacity int    `json:"capacity"`
	Current  int    `json:"current"`
	State    string `json:"state"`
	Name     string `json:"name"`
}

func (j *Jug) updateState() {
	switch j.Current {
	case 0:
		j.State = STATE_EMPTY
	case j.Capacity:
		j.State = STATE_FULL
	default:
		j.State = STATE_PARTIAL_FULL
	}
}

func (j *Jug) fill() {
	j.Current = j.Capacity
	j.updateState()
}

func (j *Jug) empty() {
	j.Current = 0
	j.updateState()
}

func (j *Jug) transferTo(target *Jug) {
	// Calculations
	amt := mathutils.Min(target.Capacity - target.Current, j.Current)
	target.Current += amt
	j.Current -= amt

	// Update state
	j.updateState()
	target.updateState()
}

func (j *Jug) dump() {
	log.Printf("Jug. Name: %s, Capacity: %d, Current: %d, State: %s\n", j.Name, j.Capacity, j.Current, j.State)
}