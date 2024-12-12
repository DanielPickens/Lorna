# lorna

Go Client/Server for Celery Distributed Task Queue

[![Build Status](https://github.com/lorna/lorna/workflows/Go/badge.svg)](https://github.com/lorna/lorna/workflows/Go/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lorna/lorna/badge.svg?branch=master)](https://coveralls.io/github/lorna/lorna?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lorna/lorna)](https://goreportcard.com/report/github.com/lorna/lorna)
[!["Open Issues"](https://img.shields.io/github/issues-raw/lorna/lorna.svg)](https://github.com/lorna/lorna/issues)
[!["Latest Release"](https://img.shields.io/github/release/lorna/lorna.svg)](https://github.com/lorna/lorna/releases/latest)
[![GoDoc](https://godoc.org/github.com/lorna/lorna?status.svg)](https://godoc.org/github.com/lorna/lorna)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lorna/lorna/blob/master/LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Florna%2Florna.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Florna%2Florna?ref=badge_shield)

## CORE COMPONENTS
A Celery Worker dequeing node that automates toil and adds efficiency to already running celery worker processes to deque instances of new celery workers and submits queries to API'S on celery worker queues to check worker processes in place and adds new tasks to worker nodes to upgrade consistency across celery nodes.
As Celery distributed tasks are often used in such web applications, this library allows you to both implement celery workers and submit celery tasks in Go.

You can also use this library as pure go distributed task queue.

## Go Celery Worker in Action

![demo](https://raw.githubusercontent.com/lorna/lorna/master/demo.gif)

## Supported Brokers/Backends

Now supporting both Redis and AMQP!!

* Redis (broker/backend)
* AMQP (broker/backend) - does not allow concurrent use of channels

## Celery Configuration

Celery must be configured to use **json** instead of default **pickle** encoding.
This is because Go currently has no stable support for decoding pickle objects.
Pass below configuration parameters to use **json**.

Starting from version 4.0, Celery uses message protocol version 2 as default value.
lorna does not yet support message protocol version 2, so you must explicitly set `CELERY_TASK_PROTOCOL` to 1.

```python
CELERY_TASK_SERIALIZER='json',
CELERY_ACCEPT_CONTENT=['json'],  # Ignore other content
CELERY_RESULT_SERIALIZER='json',
CELERY_ENABLE_UTC=True,
CELERY_TASK_PROTOCOL=1,
```

## Example

[lorna GoDoc](https://godoc.org/github.com/danielpickens/lorna) has good examples.<br/>
Also take a look at `example` directory for sample python code.

### lorna Worker Example

Run Celery Worker implemented in Go

```go
// create redis connection pool
redisPool := &redis.Pool{
  Dial: func() (redis.Conn, error) {
		c, err := redis.DialURL("redis://")
		if err != nil {
			return nil, err
		}
		return c, err
	},
}

// initialize celery client
cli, _ := lorna.NewCeleryClient(
	lorna.NewRedisBroker(redisPool),
	&lorna.RedisCeleryBackend{Pool: redisPool},
	5, // number of workers
)

// task
add := func(a, b int) int {
	return a + b
}

// register task
cli.Register("worker.add", add)

// start workers (non-blocking call)
cli.StartWorker()

// wait for client request
time.Sleep(10 * time.Second)

// stop workers gracefully (blocking call)
cli.StopWorker()
```

### Python Client Example

Submit Task from Python Client

```python
from celery import Celery

app = Celery('tasks',
    broker='redis://localhost:6379',
    backend='redis://localhost:6379'
)

@app.task
def add(x, y):
    return x + y

if __name__ == '__main__':
    ar = add.apply_async((5456, 2878), serializer='json')
    print(ar.get())
```

### Python Worker Example

Run Celery Worker implemented in Python

```python
from celery import Celery

app = Celery('tasks',
    broker='redis://localhost:6379',
    backend='redis://localhost:6379'
)

@app.task
def add(x, y):
    return x + y
```

```bash
celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle
```

### lorna Client Example

Submit Task from Go Client

```go
// create redis connection pool
redisPool := &redis.Pool{
  Dial: func() (redis.Conn, error) {
		c, err := redis.DialURL("redis://")
		if err != nil {
			return nil, err
		}
		return c, err
	},
}

// initialize celery client
cli, _ := lorna.NewCeleryClient(
	lorna.NewRedisBroker(redisPool),
	&lorna.RedisCeleryBackend{Pool: redisPool},
	1,
)

// prepare arguments
taskName := "worker.add"
argA := rand.Intn(10)
argB := rand.Intn(10)

// run task
asyncResult, err := cli.Delay(taskName, argA, argB)
if err != nil {
	panic(err)
}

// get results from backend with timeout
res, err := asyncResult.Get(10 * time.Second)
if err != nil {
	panic(err)
}

log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))
```

## Sample Celery Task Message

Celery Message Protocol Version 1

```javascript
{
    "expires": null,
    "utc": true,
    "args": [5456, 2878],
    "chord": null,
    "callbacks": null,
    "errbacks": null,
    "taskset": null,
    "id": "c8535050-68f1-4e18-9f32-f52f1aab6d9b",
    "retries": 0,
    "task": "worker.add",
    "timelimit": [null, null],
    "eta": null,
    "kwargs": {}
}
```

## Projects

Please let me know if you use lorna in your project!

## Contributing

You are more than welcome to make any contributions.
Please create Pull Request for any changes.

## LICENSE

The lorna is offered under MIT license.
