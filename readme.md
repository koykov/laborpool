# Labor pool

Workers pool implementation to reduce `runtime.findrunnable` pressure.

The main idea: reduce goroutines creation and release by storing existing goroutines to pool for further reuse.

## Usage

```go
pool := NewPool(500, 0.001)
in := make(chan interface{}, 10)
for {
	x, ok := <- in
	if !ok {
		break
	}
	worker := pool.Hire() // get existing worker from pool or create new one
	_ = worker.Do(func() error {
		// do something with x
		return nil
	})
	pool.Fire(worker) // put worker back to poll to use it later without creating
}
```
