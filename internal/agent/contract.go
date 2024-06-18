package agent

import "time"

type Agent interface {
	CollectMetrics(pollInterval time.Duration)
	SendAllMetrics(reportInterval time.Duration)
}
