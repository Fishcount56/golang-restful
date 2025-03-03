package main

import (
	"net/http"
	"restful-api/app"
)

func main() {
	db := app.NewDatabase()
	router := app.NewRouter(db)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
