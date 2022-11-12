package laborpool

// DummyMetrics is a stub metrics writer handler that uses by default and does nothing.
// Need just to reduce checks in code.
type DummyMetrics struct{}

func (DummyMetrics) Hire(bool) {}
func (DummyMetrics) Fire()     {}
func (DummyMetrics) Retire()   {}
