package main

import (
	"fmt"
	"github.com/kont1n/MetricsCollection/internal/dispatch"
	"github.com/kont1n/MetricsCollection/internal/metrics"
	"sync"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {

	var agentData []map[string]float64
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		metrics.CollectMetrics(agentData)
		time.Sleep(pollInterval)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(reportInterval)
		for pollCount, pollData := range agentData {
			if len(pollData) != 0 {
				metricType := "counter"
				metricName := "PollCount"
				err := dispatch.SendMetrics(metricType, metricName, float64(pollCount))
				if err != nil {
					fmt.Printf("failed to send counter %s: %v\n", metricName, err)
					return
				}

				metricType = "gauge"
				for metricName, metricValue := range pollData {
					err := dispatch.SendMetrics(metricType, metricName, metricValue)
					if err != nil {
						fmt.Printf("failed to send gauge %s: %v\n", metricName, err)
						return
					}
				}
			}
			clear(pollData)
		}
	}()

	wg.Wait()

}
