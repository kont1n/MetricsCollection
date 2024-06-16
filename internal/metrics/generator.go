package metrics

import (
	"math/rand"
	"runtime"
)

func CollectMetrics(agentData []map[string]float64) []map[string]float64 {

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

	agentData = append(agentData, gaugeMetric)

	return agentData
}
