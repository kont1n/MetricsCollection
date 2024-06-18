package agent

import (
	"fmt"
	"github.com/kont1n/MetricsCollection/internal/agent/api"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Measurement struct {
	Poll map[string]float64
}

type Metrics struct {
	AgentData []Measurement
	mu        sync.Mutex
	Wg        sync.WaitGroup
}

func NewAgent() *Metrics {
	metrics := Metrics{AgentData: make([]Measurement, 0)}
	return &metrics
}

func (a *Metrics) CollectMetrics(pollInterval time.Duration) {
	defer a.Wg.Done()
	//for {
	a.mu.Lock()
	measurement := CollectRuntime()
	a.AgentData = append(a.AgentData, measurement)
	a.mu.Unlock()
	time.Sleep(pollInterval)
	//}
}

func (a *Metrics) SendAllMetrics(reportInterval time.Duration) {
	defer a.Wg.Done()
	time.Sleep(reportInterval)
	fmt.Println("send", a)
	for pollCount, pollData := range a.AgentData {
		if len(pollData.Poll) != 0 {
			metricType := "gauge"
			for metricName, metricValue := range pollData.Poll {
				err := api.SendMetrics(metricType, metricName, metricValue)
				if err != nil {
					fmt.Printf("failed to send gauge %s: %v\n", metricName, err)
				}
			}

			metricType = "counter"
			metricName := "PollCount"
			err := api.SendMetrics(metricType, metricName, float64(pollCount))
			if err != nil {
				fmt.Printf("failed to send counter %s: %v\n", metricName, err)
			}
		}
		clear(pollData.Poll)
	}
}

func CollectRuntime() Measurement {

	gaugeMetric := make(map[string]float64)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	gaugeMetric["Alloc"] = float64(m.Alloc)
	gaugeMetric["BuckHashSys"] = float64(m.BuckHashSys)
	gaugeMetric["Frees"] = float64(m.Frees)
	gaugeMetric["GCCPUFraction"] = m.GCCPUFraction
	gaugeMetric["GCSys"] = float64(m.GCSys)
	gaugeMetric["HeapAlloc"] = float64(m.HeapAlloc)
	gaugeMetric["HeapIdle"] = float64(m.HeapIdle)
	gaugeMetric["HeapInuse"] = float64(m.HeapInuse)
	gaugeMetric["HeapObjects"] = float64(m.HeapObjects)
	gaugeMetric["HeapReleased"] = float64(m.HeapReleased)
	gaugeMetric["HeapSys"] = float64(m.HeapSys)
	gaugeMetric["LastGC"] = float64(m.LastGC)
	gaugeMetric["Lookups"] = float64(m.Lookups)
	gaugeMetric["MCacheInuse"] = float64(m.MCacheInuse)
	gaugeMetric["MCacheSys"] = float64(m.MCacheSys)
	gaugeMetric["MSpanInuse"] = float64(m.MSpanInuse)
	gaugeMetric["MSpanSys"] = float64(m.MSpanSys)
	gaugeMetric["Mallocs"] = float64(m.Mallocs)
	gaugeMetric["NextGC"] = float64(m.NextGC)
	gaugeMetric["NumForcedGC"] = float64(m.NumForcedGC)
	gaugeMetric["NumGC"] = float64(m.NumGC)
	gaugeMetric["OtherSys"] = float64(m.OtherSys)
	gaugeMetric["PauseTotalNs"] = float64(m.PauseTotalNs)
	gaugeMetric["StackInuse"] = float64(m.StackInuse)
	gaugeMetric["StackSys"] = float64(m.StackSys)
	gaugeMetric["Sys"] = float64(m.Sys)
	gaugeMetric["TotalAlloc"] = float64(m.TotalAlloc)
	gaugeMetric["RandomValue"] = rand.Float64()

	measurement := Measurement{Poll: gaugeMetric}
	return measurement
}
