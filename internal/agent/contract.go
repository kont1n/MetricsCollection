package agent

import "time"

type Agent interface {
	CollectMetrics(pollInterval time.Duration) AgentMetrics
	SendAllMetrics(reportInterval time.Duration) error
}
