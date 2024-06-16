package api

import (
	"github.com/kont1n/MetricsCollection/internal/storage"
	"net/http"
	"strconv"
)

const (
	ContentTypeText = "text/plain"
)

func UpdateHandler(s storage.MetricStorage) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(res, "Only POST requests are allowed!", http.StatusMethodNotAllowed)
			return
		}

		metricType := req.PathValue("metricType")
		metricName := req.PathValue("metricName")
		switch metricType {
		case "gauge":
			{
				if value, err := strconv.ParseFloat(req.PathValue("metricValue"), 64); err == nil {
					s.ValueReplace(metricType, metricName, value)
				} else {
					http.Error(res, err.Error(), http.StatusBadRequest)
				}
			}
		case "counter":
			{
				if value, err := strconv.ParseInt(req.PathValue("metricValue"), 10, 64); err == nil {
					s.ValueIncrement(metricType, metricName, value)
				} else {
					http.Error(res, err.Error(), http.StatusBadRequest)
				}
			}
		default:
			{
				http.Error(res, "Metric type not found!", http.StatusBadRequest)
				return
			}
		}
		res.Header().Set("Content-Type", ContentTypeText)
		res.WriteHeader(http.StatusOK)
	}
}
