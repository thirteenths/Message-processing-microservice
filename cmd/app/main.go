package main

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/thirteenths/message-processing-microservice/internal/app"
	"github.com/thirteenths/message-processing-microservice/internal/ports/handlers"
	"github.com/thirteenths/message-processing-microservice/internal/storage/postgres"
)

// @title           Message processing microservice
// @version         1.0
// @description     This is backend server for test task 2024.

// @host      localhost:3000
// @BasePath  /api/

func main() {
	storage, err := postgres.NewMessageStorage(os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal(err)
	}

	messService := app.NewMessageService(storage)

	handler := handlers.NewApiHandler(*messService)

	r := chi.NewRouter()
	r.Post("/api/message", handler.CreateMessage)
	r.Get("/api/statistic", handler.GetStatistic)

	r.Get("/api/docs/*", httpSwagger.WrapHandler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
