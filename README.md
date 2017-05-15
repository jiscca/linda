# Linda

Linda is a dispatcher and scheduler system.It allows you to save tasks and push a job of this task into a queue.

Linda is based on redis and cassandra, and it will support more backend.

# Installation

To install Linda, use 

`go get github.com/amlun/linda`

to install the package, and then use

`glide install`

to install the dependency packages

# Getting Started

## Simple Usage

Edit the apps/server.go, apps/scheduler.go, modify the config with your own urlString

```go
package main

import (
	"github.com/amlun/linda/linda"
	"github.com/amlun/linda/modules/server"
)

func main() {
	var config = linda.Config{
		BrokerURL: "redis://127.0.0.1:6379",
		SaverURL:  "cassandra://cassandra:cassandra@127.0.0.1:9042/linda",
	}
	l := linda.NewLinda(&config)
	server.Start(l)
}
```

```go
package main

import (
	"github.com/amlun/linda/linda"
	"github.com/amlun/linda/modules/scheduler"
)

func main() {
	var config = linda.Config{
		BrokerURL: "redis://127.0.0.1:6379",
		SaverURL:  "cassandra://cassandra:cassandra@127.0.0.1:9042/linda",
	}
	l := linda.NewLinda(&config)
	s := scheduler.New(l)
	s.Start()
}

```

And then use

`go run apps/server.go`

to start a http server and serve the apis.

Use

`go run apps/scheduler.go`

to start scheduler to schedule the timing task.


## Api Doc

 * GET /api/ping - Check the server if it is alive.
 * GET /api/tasks - List all tasks.
 * GET /api/queues - List all queues and its length.
 * GET /api/job - Get a job from queue, now it only implements a simple way to fetch a job.
 * POST /api/task - Post a task.