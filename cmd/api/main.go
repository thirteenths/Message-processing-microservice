package main

import (
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/thirteenths/message-processing-microservice/cmd/api/docs"

	"net/http"
	"os"
	"time"

	"github.com/thirteenths/message-processing-microservice/internal/app"
	"github.com/thirteenths/message-processing-microservice/internal/ports/https"
	"github.com/thirteenths/message-processing-microservice/internal/storage"
	"github.com/thirteenths/message-processing-microservice/internal/storage/kafka"
	"github.com/thirteenths/message-processing-microservice/internal/storage/postgres"
)

// @title           Message processing microservice
// @version         1.0
// @description     This is backend server for test task 2024.

// @host      localhost:3000
// @BasePath  /api/

func main() {
	pg, err := postgres.NewMessageStorage(os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to PostgreSQL")

	k, err := kafka.NewKafkaClient("localhost:9092", "topic", 0)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to Kafka")

	stg := storage.NewStorage(pg, k)
	msgService := app.NewMessageService(stg)
	handler := https.NewApiHandler(*msgService)

	r := chi.NewRouter()
	r.Post("/api/message", handler.CreateMessage)
	r.Get("/api/statistic", handler.GetStatistic)
	r.Get("/api/message", handler.GetMessage)

	r.Get("/api/docs/*", httpSwagger.WrapHandler)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
