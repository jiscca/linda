package scheduler

import (
	"github.com/amlun/linda/linda"
	cron "github.com/carlescere/scheduler"
	"runtime"
)

type scheduler struct {
	linda    *linda.Linda
	isMaster bool
}

func New(linda *linda.Linda) *scheduler {
	return &scheduler{
		linda: linda,
	}
}

// will support distribute deploy
// use redis or zookeeper to lock ,one master with multi slave
func (s *scheduler) Start() {
	list := s.linda.Schedules()
	for _, frequency := range list {
		cron.Every(int(frequency.Seconds())).Seconds().Run(s.linda.Schedule(frequency))
	}
	runtime.Goexit()
}