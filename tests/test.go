package main

import (
	"testing"
	"time"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func TestTelemetryResultStatus(t *testing.T) {
	res := telemetry.Result{
		Timestamp: time.Now(),
		Target:    "127.0.0.1",
		Layer:     "L3",
		Probe:     "icmp",
		Status:    telemetry.StatusUp,
		Latency:   time.Millisecond,
	}

	if res.Status != telemetry.StatusUp {
		t.Errorf("expected status UP, got %s", res.Status)
	}
}
