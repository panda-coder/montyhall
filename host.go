package main

import (
	"math/rand"
	"time"
)

type Host struct {
	gates [3]bool
	rng   *rand.Rand
}

func (h *Host) Game() {

}

func (h *Host) GetGate() int {
	for idx, _ := range h.gates {
		if h.gates[idx] {
			return idx
		}
	}
	return -1
}

func (h *Host) SetupGate() {
	for idx, _ := range h.gates {
		h.gates[idx] = false
	}
	n := h.rng.Intn(totalGates)
	h.gates[n] = true
}

func NewHost() *Host {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &Host{
		rng: rng,
	}
}
