package linda

import "time"

type Config struct {
	Queue     string
	Timeout   int64
	Interval  time.Duration
	PollerNum int
	WorkerNum int
}
