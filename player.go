package main

import (
	"math/rand"
	"time"
)

type Player struct {
	maintain     bool
	selectedGate int
	rng          *rand.Rand
}

func (p *Player) SelectGate() int {
	n := p.rng.Intn(totalGates)
	p.selectedGate = n

	return p.selectedGate
}

func (p *Player) GetMaintain() bool {
	return p.maintain
}

func NewPlayer(maintain bool) *Player {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return &Player{
		rng:      rng,
		maintain: maintain,
	}
}
