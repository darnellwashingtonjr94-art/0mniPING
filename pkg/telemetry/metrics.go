package telemetry

import (
	"time"
)

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type Result struct {
	Timestamp time.Time     `json:"timestamp"`
	Target    string        `json:"target"`
	Layer     string        `json:"layer"`
	Probe     string        `json:"probe"`
	Status    Status        `json:"status"`
	Latency   time.Duration `json:"latency"`
	Error     string        `json:"error,omitempty"`
}
