package main

import (
	"log"
	"net/http"
	"time"

	"github.com/KKolyasik/observability-service/internal/handler"
	"github.com/KKolyasik/observability-service/internal/service"
	"github.com/KKolyasik/observability-service/internal/storage"
)

func main() {
	st := storage.NewMemStorage()
	svc := service.New(st)

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handler.UpdateMetrics(svc))

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
