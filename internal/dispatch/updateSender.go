package dispatch

import (
	"io"
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

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", contentTypeText)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	_, err = io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return nil
}
