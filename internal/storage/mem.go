package storage


type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (m *MemStorage) GetGauge(name string) (float64, bool) {
	v, ok := m.gauge[name]
	return v, ok
}

func (m *MemStorage) GetCounter(name string) (int64, bool) {
	v, ok := m.counter[name]
	return v, ok
}

func (m *MemStorage) SetGauge(name string, value float64) error {
	m.gauge[name] = value
	return nil
}

func (m *MemStorage) IncCounter(name string, value int64) error {
	m.counter[name] += value
	return nil
}
