package laborpool

type DummyMetrics struct{}

func (DummyMetrics) Hire(bool) {}
func (DummyMetrics) Fire()     {}
func (DummyMetrics) Retire()   {}
