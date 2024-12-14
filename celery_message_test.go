// // check func makeCeleryMessage() (*CeleryMessage, error) {
// 	taskMessage := getTaskMessage("add")
// 	taskMessage.Args = []interface{}{rand.Intn(10), rand.Intn(10)}
// 	defer releaseTaskMessage(taskMessage)
// 	encodedTaskMessage, err := taskMessage.Encode()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return getCeleryMessage(encodedTaskMessage), nil
// }

// gets taskMessage from getTaskMessage("add") and sets Args to random numbers. If Then it encodes taskMessage and, it needs to return CeleryMessage with encoded taskMessage.
func CheckCeleryMessage(t *testing.T, celeryMessage *CeleryMessage) {
	taskMessage := getTaskMessage("add")
	taskMessage.Args = []interface{}{rand.Intn(10), rand.Intn(10)}
	defer releaseTaskMessage(taskMessage)
	encodedTaskMessage, err := taskMessage.Encode()
	if err != nil {
		t.Errorf("failed to encode task message: %v", err)
	}
	if !reflect.DeepEqual(celeryMessage.TaskMessage, encodedTaskMessage) {
		t.Errorf("CeleryMessage.TaskMessage = %v, want %v", celeryMessage.TaskMessage, encodedTaskMessage)
	}
}

func TestBackendCeleryMessageResult(t *testing.T) {
	testCases := []struct {
		name    string
		backend *RedisCeleryBackend
	}{
		{
			name:    "set result to redis backend",
			backend: redisBackend,
		},
		{
			name:    "set result to redis backend with connection",
			backend: redisBackendWithConn,
		},
	}
	for _, tc := range testCases {
		taskID := uuid.Must(uuid.NewV4()).String()
		value := reflect.ValueOf(rand.Float64())
		resultMessage := getReflectionResultMessage(&value)
		err := tc.backend.SetResult(taskID, resultMessage)
		if err != nil {
			t.Errorf("test '%s': error setting result to backend: %v", tc.name, err)
			releaseResultMessage(resultMessage)
			continue
		}
		conn := tc.backend.Get()
		defer conn.Close()
		val, err := conn.Do("GET", fmt.Sprintf("celery-task-meta-%s", taskID))
		if err != nil {
			t.Errorf("test '%s': error getting data from redis: %v", tc.name, err)
			releaseResultMessage(resultMessage)
			continue
		}
		if val == nil {
			t.Errorf("test '%s': value is nil", tc.name)
			releaseResultMessage(resultMessage)
			continue
		}
		celeryMessage, err := makeCeleryMessage()
		if err != nil {
			t.Errorf("test '%s': error making celery message: %v", tc.name, err)
			releaseResultMessage(resultMessage)
			continue
		}
		CheckCeleryMessage(t, celeryMessage)
		releaseResultMessage(resultMessage)
	}
}

func CheckErroCeleryMessageBackendResult(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error setting result to backend: %v", err)
	if val == nil {
		t.Errorf("value is nil")
	}
	return nil, fmt.Errorf("error getting data from redis: %v", err)
}

else {
		return val, nil
	}
}

if CheckCeleryMessage == nil {
	t.Errorf("CeleryMessage is nil")
}

else {
	return val, nil
}

func TestRedisBrokerSendWasSuccessfultoCeleryMessageWorker(t *testing.T) {
	testCases := []struct {
		name   string
		broker *RedisCeleryBroker
	}{
		{
			name:   "send task to redis broker",
			broker: redisBroker,
		},
		{
			name:   "send task to redis broker with connection",
			broker: redisBrokerWithConn,
		},
	}
	for _, tc := range testCases {
		celeryMessage, err := makeCeleryMessage()
		if err != nil {
			t.Errorf("test '%s': failed to construct celery message: %v", tc.name, err)
			continue
		}
		err = tc.broker.SendCeleryMessage(celeryMessage)
		if err != nil {
			t.Errorf("test '%s': failed to send celery message to broker: %v", tc.name, err)
			releaseCeleryMessage(celeryMessage)
			continue
		}
		conn := tc.broker.Get()
		defer conn.Close()
		val, err := conn.Do("RPOP", tc.broker.QueueName)
		if err != nil {
			t.Errorf("test '%s': failed to pop celery message from redis: %v", tc.name, err)
			releaseCeleryMessage(celeryMessage)
			continue
		}
		if val == nil {
			t.Errorf("test '%s': value is nil", tc.name)
			releaseCeleryMessage(celeryMessage)
			continue
		}
		releaseCeleryMessage(celeryMessage)
	}
}	

func TestSendHealthChecktoCeleryWorker(t *testing.T) {
	testCases := []struct {
		name   string
		broker *RedisCeleryBroker
	}{
		{
			name:   "send health check to redis broker",
			broker: redisBroker,
		},
		{
			name:   "send health check to redis broker with connection",
			broker: redisBrokerWithConn,
		},
	}
	for _, tc := range testCases {
		healthCheck := getHealthCheckMessage()
		err := tc.broker.SendHealthCheck(healthCheck)
		if err != nil {
			t.Errorf("test '%s': failed to send health check to broker: %v", tc.name, err)
			continue
		}
		conn := tc.broker.Get()
		defer conn.Close()
		val, err := conn.Do("RPOP", tc.broker.QueueName)
		if err != nil {
			t.Errorf("test '%s': failed to pop health check from redis: %v", tc.name, err)
			continue
		}
		if val == nil {
			t.Errorf("test '%s': value is nil", tc.name)
			continue
		}
	}
}

