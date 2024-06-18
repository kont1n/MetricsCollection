package main

import (
	"github.com/kont1n/MetricsCollection/internal/agent"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {

	a := agent.NewAgent()

	a.Wg.Add(2)
	a.CollectMetrics(pollInterval)
	a.SendAllMetrics(reportInterval)
	a.Wg.Wait()

}
