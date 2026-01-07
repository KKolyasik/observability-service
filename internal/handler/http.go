package handler

import (
	"net/http"
	"strings"
)

type MetricsUpdater interface {
	Update(metricType, metricName, metricValue string) error
}

func UpdateMetrics(svc MetricsUpdater) http.HandlerFunc {
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

		metricType, metricName, metricValue := parts[0], parts[1], parts[2]

		err := svc.Update(metricType, metricName, metricValue)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusOK)
	}
}
