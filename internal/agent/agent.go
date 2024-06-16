package agent

import (
	"fmt"
	"github.com/kont1n/MetricsCollection/internal/agent/api"
	"github.com/kont1n/MetricsCollection/internal/agent/metrics"
	"time"
)

type AgentMetrics struct {
	agentData []map[string]float64
}

func NewAgent() AgentMetrics {
	return AgentMetrics{agentData: []map[string]float64{}}
}

func (a *AgentMetrics) CollectMetrics(pollInterval time.Duration) AgentMetrics {
	metrics.Collect(a.agentData)
	time.Sleep(pollInterval)
	return *a
}

func (a *AgentMetrics) SendAllMetrics(reportInterval time.Duration) error {
	time.Sleep(reportInterval)
	for pollCount, pollData := range a.agentData {
		if len(pollData) != 0 {
			metricType := "gauge"
			for metricName, metricValue := range pollData {
				err := api.SendMetrics(metricType, metricName, metricValue)
				if err != nil {
					fmt.Printf("failed to send gauge %s: %v\n", metricName, err)
					return err
				}
			}

			metricType = "counter"
			metricName := "PollCount"
			err := api.SendMetrics(metricType, metricName, float64(pollCount))
			if err != nil {
				fmt.Printf("failed to send counter %s: %v\n", metricName, err)
				return err
			}
		}
		clear(pollData)
	}
	return nil
}
