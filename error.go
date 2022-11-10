package laborpool

import "errors"

var (
	ErrWorkerStatus = errors.New("uninitialized worker, use pool.Hire()")
)
