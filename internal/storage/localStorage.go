package storage

type Metric struct {
	metricType  string
	metricName  string
	metricValue float64
}

type MetricStorage struct {
	metricData map[string]Metric
}

func CreateLocalStorage() MetricStorage {
	return MetricStorage{metricData: make(map[string]Metric)}
}

func (s MetricStorage) ValueReplace(metricType string, metricName string, value float64) {
	s.metricData[(metricType + metricName)] = Metric{metricType, metricName, value}
}

func (s MetricStorage) ValueIncrement(metricType string, metricName string, value int64) {
	var m, ok = s.metricData[(metricType + metricName)]
	if ok {
		s.metricData[metricType+metricName] = Metric{metricType, metricName, m.metricValue + float64(value)}
	} else {
		s.metricData[metricType+metricName] = Metric{metricType, metricName, float64(value)}
	}
}
