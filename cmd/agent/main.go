package main

import (
	"fmt"
	"github.com/kont1n/MetricsCollection/internal/agent"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {

	a := agent.NewAgent()
	a.CollectMetrics(pollInterval)
	err := a.SendAllMetrics(reportInterval)
	if err != nil {
		fmt.Printf("failed to send metrics: %v\n", err)
	}
}
