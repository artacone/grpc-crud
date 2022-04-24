package generator

import "sync"

type Generator interface {
	Next() uint64
}

type uintGenerator struct {
	n    uint64
	step uint64
	mu   sync.Mutex
}

func New(start, step uint64) Generator {
	return &uintGenerator{
		n:    start,
		step: step,
	}
}

func (g *uintGenerator) Next() uint64 {
	g.mu.Lock()
	n := g.n
	g.n += g.step
	g.mu.Unlock()
	return n
}
