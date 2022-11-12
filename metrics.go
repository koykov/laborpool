package laborpool

// MetricsWriter is an interface of metrics handler.
// See example of implementations https://github.com/koykov/metrics_writers/tree/master/laborpool.
type MetricsWriter interface {
	// Hire registers taking worker from the pool.
	Hire(unknown bool)
	// Fire registers returning worker to the pool.
	Fire()
	// Retire registers worker leak (with release).
	Retire()
}
