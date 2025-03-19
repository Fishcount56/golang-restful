package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restful-api/dto/response"
	"restful-api/queue"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/julienschmidt/httprouter"
)

var redisClient = asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

func HealthController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.NewApiResponseBuilder().SetStatus(200).SetMessage("OK").SetRequestId(uuid.New().String()).Build()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func EnqueueTaskController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ids := ps.ByName("id")

	payload, err := json.Marshal(map[string]string{
		"id": ids,
	})

	if err != nil {
		http.Error(w, "Failed to create payload", http.StatusInternalServerError)
		return
	}

	task := asynq.NewTask(queue.TestingQueue, payload, asynq.MaxRetry(3), asynq.Timeout(30*time.Second))

	info, err := redisClient.Enqueue(task, asynq.Queue("testingQueue"))
	if err != nil {
		log.Println("Failed to enqueue job:", err)
		http.Error(w, "Failed to enqueue job", http.StatusInternalServerError)
	}

	res := response.NewApiResponseBuilder().
		SetStatus(200).
		SetMessage("Job added to queue").
		SetRequestId(info.ID).
		Build()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
