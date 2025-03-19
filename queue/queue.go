package queue

import "github.com/hibiken/asynq"

const TestingQueue = "testingQueue"

func NewClient() *asynq.Client {
	redisOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	return asynq.NewClient(redisOpt)
}

func NewServer() *asynq.Server {
	redisOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	return asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			TestingQueue: 1,
		},
	})
}
