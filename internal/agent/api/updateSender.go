package api

import (
	"net/http"
	"strconv"
)

const (
	contentTypeText = "text/plain"
	serverProtocol  = "http"
	serverAddress   = "127.0.0.1"
	serverPort      = "8080"
	serverPath      = "update"
)

func SendMetrics(metricType string, metricName string, metricValue float64) error {
	var url = serverProtocol + "://" + serverAddress + ":" + serverPort + "/" + serverPath + "/" + metricType + "/" + metricName + "/" + strconv.FormatFloat(metricValue, 'f', -1, 64)

	res, err := http.Post(url, contentTypeText, nil)
	if err != nil {
		return err
	}

	err = res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
