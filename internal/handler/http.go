package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/KKolyasik/observability-service/internal/models"
)

type MetricsService interface {
    UpdateGauge(name string, value float64) error
    IncCounter(name string, delta int64) error
}


func UpdateMetrics(svc MetricsService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Only method POST is allowed", http.StatusMethodNotAllowed)
            return
        }

        path := strings.TrimPrefix(r.URL.Path, "/update/")
        path = strings.Trim(path, "/")
        if path == "" {
            http.NotFound(w, r)
            return
        }

        parts := strings.Split(path, "/")

        if len(parts) < 3 {
            http.NotFound(w, r)
            return
        }
        if len(parts) > 3 {
            http.Error(w, "invalid path format, expected /update/{type}/{name}/{value}", http.StatusBadRequest)
            return
        }

        metricType, metricName, raw := parts[0], parts[1], parts[2]
        if metricName == "" {
            http.NotFound(w, r)
            return
        }

        switch metricType {
        case models.Gauge:
            v, err := strconv.ParseFloat(raw, 64)
            if err != nil {
                http.Error(w, "invalid gauge value", http.StatusBadRequest)
                return
            }
            if err := svc.UpdateGauge(metricName, v); err != nil {
                http.Error(w, "failed to store metric", http.StatusInternalServerError)
                return
            }

        case models.Counter:
            v, err := strconv.ParseInt(raw, 10, 64)
            if err != nil {
                http.Error(w, "invalid counter value", http.StatusBadRequest)
                return
            }
            if err := svc.IncCounter(metricName, v); err != nil {
                http.Error(w, "failed to store metric", http.StatusInternalServerError)
                return
            }

        default:
            http.Error(w, "unknown metric type (use gauge|counter)", http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}
