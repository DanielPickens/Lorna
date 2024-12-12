// Copyright (c) 2024 Daniel Pickens
// This file is part of lorna which is released under MIT license.
// See file LICENSE for full license details.

package lorna

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func Example_worker() {

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
	cli, _ := NewCeleryClient(
		NewRedisBroker(redisPool),
		&RedisCeleryBackend{Pool: redisPool},
		5, // number of workers
	)

	// task
	add := func(a, b int) int {
		return a + b
	}

	// register task
	cli.Register("add", add)

	// start workers (non-blocking call)
	cli.StartWorker()

	// wait for client request
	time.Sleep(10 * time.Second)

	// stop workers gracefully (blocking call)
	cli.StopWorker()

}
