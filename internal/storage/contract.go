package storage

type Storage interface {
	ValueReplace(metricType, metricName string, value float64)
	ValueIncrement(metricType, metricName string, value int64)
}
