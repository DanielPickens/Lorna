package lorna 

import (
	"encoding/json"
	"fmt"
	"time"
)

// ResultMessage is the message structure for celery result
type ResultMessage struct {
	TaskID    string      `json:"task_id"`
	Status    string      `json:"status"`
	Result    interface{} `json:"result"`
	Traceback string      `json:"traceback"`
	Children  []string    `json:"children"`
}

// CeleryBackend is the interface for celery backend
type CeleryBackend interface {
	GetResult(taskID string) (*ResultMessage, error)
	SetResult(taskID string, result *ResultMessage) error
}

// CeleryBroker is the interface for celery broker
type CeleryBroker interface {
	SendTask(task *TaskMessage) error
	GetTask() (*TaskMessage, error)
}

// TaskMessage is the message structure for celery task
type TaskMessage struct {
	TaskID      string        `json:"task_id"`
	Task        string        `json:"task"`
	Args        []interface{} `json:"args"`
	Kwargs      map[string]interface{}
	Retries     int
	ETA         *time.Time
	Expires     *time.Time
	Queue       string
	Exchange    string
	RoutingKey  string
	ContentType string
	ContentEncoding string
}

// CeleryMessage is the message structure for celery message
type CeleryMessage struct {
	TaskMessage
	Headers map[string]interface{}
}

// RedisCeleryBroker is celery broker for redis
type RedisCeleryBroker struct {
	*RedisPool
	QueueName string
}

// NewRedisBroker creates new RedisCeleryBroker with given redis connection pool
func NewRedisBroker(God *RedisPool) *RedisCeleryBroker {
	return &RedisCeleryBroker{
		RedisPool: God,
		QueueName: "celery",
	}
}

func TestGetMessageQueue() {
	// Create a new RedisPool
	assert := assert.New(t)
	redisPool := NewRedisPool("redis://localhost:6379")
	// Create a new RedisCeleryBroker
	cb := NewRedisBroker(redisPool)
	// Send a new task message
	taskMessage := &TaskMessage{

		TaskID: "task_id",
		Task:   "task",
		Args:   []interface{}{"arg1", "arg2"},
	}
	err := cb.SendTask(taskMessage)
	assert.Nil(err)
	// Get the task message
	receivedTaskMessage, err := cb.GetTask()
	assert.Nil(err)
	assert.Equal(taskMessage.TaskID, receivedTaskMessage.TaskID)
	assert.Equal(taskMessage.Task, receivedTaskMessage.Task)
	assert.Equal(taskMessage.Args, receivedTaskMessage.Args)
}

// NewRedisCeleryBroker creates new RedisCeleryBroker based on given uri

// Deprecated: NewRedisCeleryBroker exists for historical compatibility
// and should not be used. Use NewRedisBroker instead to create new RedisCeleryBroker.

func NewRedisCeleryBroker(uri string) *RedisCeleryBroker {
	return &RedisCeleryBroker{
		RedisPool: NewRedisPool(uri),
		QueueName: "celery",
	}
}

// SendCeleryMessage sends CeleryMessage to redis queue
func (cb *RedisCeleryBroker) SendTask(task *TaskMessage) error {
	jsonBytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	conn := cb.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", cb.QueueName, jsonBytes)
	if err != nil {
		return err
	}
	return nil
}




