package laborpool

import (
	"sync"

	"github.com/koykov/lbpool"
)

type Pool struct {
	// Size indicates how many workers may be stored in the pool.
	// Workers that overflows that limit will release immediately.
	Size uint
	// PensionFactor indicates the possibility to retire worker to the pension.
	PensionFactor float32
	// MetricsWriter writers pool metrics.
	MetricsWriter MetricsWriter

	// p is an underlying lbpool instance that stores workers.
	p *lbpool.Pool
	// CC of MetricsWriters.
	mw MetricsWriter

	once sync.Once
}

// NewPool makes new pool instance with dummy MetricsWriter.
func NewPool(size uint, rate float32) *Pool {
	return NewPoolWM(size, rate, nil)
}

// NewPoolWM makes new pool instance with given MetricsWriter.
func NewPoolWM(size uint, rate float32, mw MetricsWriter) *Pool {
	p := Pool{
		Size:          size,
		PensionFactor: rate,
		MetricsWriter: mw,
	}
	return &p
}

// Internal pool initialization.
func (p *Pool) init() {
	p.p = lbpool.NewPool(p.Size, p.PensionFactor)
	if p.MetricsWriter == nil {
		p.MetricsWriter = DummyMetrics{}
	}
	p.mw = p.MetricsWriter
}

// Hire returns new worker from the underlying pool.
func (p *Pool) Hire() *Worker {
	p.once.Do(p.init)
	raw := p.p.Get()
	var unk bool
	if raw == nil {
		// No workers found un a pool, then train new one and mark it as unknown.
		unk = true
		raw = trainWorker()
	}
	if w, ok := raw.(*Worker); ok {
		p.mw.Hire(unk)
		return w
	}
	return nil
}

func (p *Pool) Fire(w *Worker) {
	if ok := p.p.Put(w); ok {
		p.mw.Fire()
		return
	}
	p.mw.Retire()
}

var _ = NewPool
