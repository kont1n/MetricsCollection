package main

import (
	"github.com/kont1n/MetricsCollection/internal/receive"
	"github.com/kont1n/MetricsCollection/internal/storage"
	"net/http"
)

func main() {

	memStorage := storage.CreateLocalStorage()

	mux := http.NewServeMux()
	mux.HandleFunc(`/update/{metricType}/{metricName}/{metricValue}`, receive.UpdateHandler(memStorage))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
