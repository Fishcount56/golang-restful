package main

import (
	"log"
	"net/http"
	"restful-api/app"
	"restful-api/queue"
	"restful-api/queue/worker"

	"github.com/hibiken/asynq"
)

func main() {
	db := app.NewDatabase()
	router := app.NewRouter(db)

	// Inisialisasi worker dari queue package
	server := queue.NewServer()

	// Inisialisasi handler worker
	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.TestingQueue, worker.HandleTestingWorker)

	// Jalankan worker dalam goroutine
	go func() {
		log.Println("Worker berjalan...")
		if err := server.Run(mux); err != nil {
			log.Fatal(err)
		}
	}()

	// Jalankan HTTP server
	serverHTTP := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	log.Println("HTTP server berjalan di http://localhost:3000")
	if err := serverHTTP.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
