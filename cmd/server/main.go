package main

import (
	"net/http"
	"github.com/KKolyasik/observability-service/internal/models"
	"github.com/KKolyasik/observability-service/internal/handler"
	"github.com/KKolyasik/observability-service/internal/service"
)



func main() {
	storage := models.NewMemStorage()
	svc := service.New(storage)

	mux := http.NewServeMux()
	mux.HandleFunc("/update/", handler.UpdateMetrics(svc))
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
