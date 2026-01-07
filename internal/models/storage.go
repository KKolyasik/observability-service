package models

type MetricsStorage interface {
    GetGauge(name string) (float64, bool)
    GetCounter(name string) (int64, bool)

    SetGauge(name string, value float64) error
    IncCounter(name string, delta int64) error
}
