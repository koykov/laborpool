package laborpool

type MetricsWriter interface {
	Hire(unknown bool)
	Fire()
	Retire()
}
