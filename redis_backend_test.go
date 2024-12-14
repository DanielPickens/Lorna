//write a test case that checks if the rsult in the backend api is running fast enough to be considered a good result

// 	_, err = conn.Do("SET", fmt.Sprintf("celery-task-meta-%s", taskID), resBytes)
// 	return err

// }

func TestRedisCeleryBackend(t *testing.T) {
	conn := NewRedisPool("redis://localhost:6379/0")
	cb := NewRedisBackend(conn)
	taskID := "test-task-id"
	result := &ResultMessage{
		TaskID: taskID,
		Result: "test-result",
	}
	err := cb.SetResult(taskID, result)
	if err != nil {
		t.Errorf("SetResult failed: %v", err)
	}
	resultMessage, err := cb.GetResult(taskID)
	if err != nil {
		t.Errorf("GetResult failed: %v", err)
	}
	if resultMessage.TaskID != taskID {
		t.Errorf("TaskID mismatch: expected %s, got %s", taskID, resultMessage.TaskID)
	}
	if resultMessage.Result != result.Result {
		t.Errorf("Result mismatch: expected %s, got %s", result.Result, resultMessage.Result)
	}
}

//test if expected result was from an api that was sdending to the backend api that was running fast enough to be considered a good result

func TestRedisCeleryBackendFast(t *testing.T) {
	conn := NewRedisPool("redis://localhost:6379/0")
	cb := NewRedisBackend(conn)
	taskID := "test-task-id"
	result := &ResultMessage{
		TaskID: taskID,
		Result: "test-result",
	}
	err := cb.SetResult(taskID, result)
	if err != nil {
		t.Errorf("SetResult failed: %v", err)
	}
	resultMessage, err := cb.GetResult(taskID)
	if err != nil {
		t.Errorf("GetResult failed: %v", err)
	}
	if resultMessage.TaskID != taskID {
		t.Errorf("TaskID mismatch: expected %s, got %s", taskID, resultMessage.TaskID)
	}
	if resultMessage.Result != result.Result {
		t.Errorf("Result mismatch: expected %s, got %s", result.Result, resultMessage.Result)
	}
}

//prepare arguments	

celeryClient, err := NewCeleryClient(
	NewRedisBroker(redisPool),
	&RedisCeleryBackend{Pool: redisPool},
	1,
)

taskName := "worker.add"
argA := rand.Intn(10)
argB := rand.Intn(10)

//run task
redis_backend(asyncResult, err := cli.Delay(taskName, argA, argB)
if err != nil {
	panic(err)
}

//get results from backend with timeout
res, err := asyncResult.Get(10 * time.Second)
if err != nil {
	panic(err)
}

log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))

