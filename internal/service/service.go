package service

import (
	"strconv"
	"errors"

	"github.com/KKolyasik/observability-service/internal/models"
)

var ErrUnknownMetricType = errors.New("unknown metric type")

type MetricsService struct {
	st models.MetricsStorage
}

func New(st models.MetricsStorage) *MetricsService {
	return &MetricsService{st: st}
}

func (m *MetricsService) Update(metricType, metricName, metricValue string) error {
	switch metricType {
	case "gauge":
		if val, err := strconv.ParseFloat(metricValue, 64); err != nil {
			return err
		} else {
			m.st.SetGauge(metricName, val)
			return nil
		}
	case "counter":
		if val, err := strconv.ParseInt(metricValue, 10, 64); err != nil {
			return err
		} else {
			m.st.IncCounter(metricName, val)
			return nil
		}
	default:
		return ErrUnknownMetricType
	}
}
