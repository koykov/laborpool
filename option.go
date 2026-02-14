package laborpool

type Option func(p *pool)

func WithMetricsWriter(mw MetricsWriter) Option {
	return func(p *pool) {
		p.mw = mw
	}
}
