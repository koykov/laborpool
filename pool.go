package laborpool

import (
	"sync"

	"github.com/koykov/lbpool"
)

type Pool struct {
	Size          uint
	PensionRate   float32
	MetricsWriter MetricsWriter

	p    *lbpool.Pool
	mw   MetricsWriter
	once sync.Once
}

func NewPool(size uint, rate float32) *Pool {
	return NewPoolWM(size, rate, nil)
}

func NewPoolWM(size uint, rate float32, mw MetricsWriter) *Pool {
	p := Pool{
		Size:          size,
		PensionRate:   rate,
		MetricsWriter: mw,
	}
	return &p
}

func (p *Pool) init() {
	p.p = lbpool.NewPool(p.Size, p.PensionRate)
	if p.MetricsWriter == nil {
		p.MetricsWriter = DummyMetrics{}
	}
	p.mw = p.MetricsWriter
}

func (p *Pool) Hire() *Worker {
	p.once.Do(p.init)
	raw := p.p.Get()
	var unk bool
	if raw == nil {
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
