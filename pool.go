package laborpool

import (
	"github.com/koykov/lbpool"
)

type Pool interface {
	Hire() Worker
	Fire(w Worker)
}

type pool struct {
	// size indicates how many workers may be stored in the pool.
	// Workers that overflows that limit will release immediately.
	size uint
	// rate indicates the possibility to retire worker to the pension.
	rate float32

	// p is an underlying lbpool instance that stores workers.
	p *lbpool.Pool
	// CC of MetricsWriters.
	mw MetricsWriter
}

// NewPool makes new pool instance with dummy MetricsWriter.
func NewPool(size uint, rate float32, options ...Option) Pool {
	p := &pool{
		size: size,
		rate: rate,
	}
	for _, fn := range options {
		fn(p)
	}
	if p.mw == nil {
		p.mw = dummyMW{}
	}
	return p
}

// Hire returns new worker from the underlying pool.
func (p *pool) Hire() Worker {
	raw := p.p.Get()
	var unk bool
	if raw == nil {
		// No workers found un a pool, then train new one and mark it as unknown.
		unk = true
		raw = trainWorker()
	}
	if w, ok := raw.(*worker); ok {
		p.mw.Hire(unk)
		return w
	}
	return nil
}

func (p *pool) Fire(w Worker) {
	if ok := p.p.Put(w); ok {
		p.mw.Fire()
		return
	}
	p.mw.Retire()
}

var _ = NewPool
