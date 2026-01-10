package service

import (
	"errors"

	"github.com/KKolyasik/observability-service/internal/storage"
)

var ErrUnknownMetricType = errors.New("unknown metric type")

type MetricsService struct {
	st storage.MetricsStorage
}

func New(st storage.MetricsStorage) *MetricsService {
	return &MetricsService{st: st}
}

func (s *MetricsService) UpdateGauge(name string, value float64) error {
	return s.st.SetGauge(name, value)
}

func (s *MetricsService) IncCounter(name string, delta int64) error {
	return s.st.IncCounter(name, delta)
}
