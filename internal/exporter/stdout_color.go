package exporter

import (
	"fmt"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
)

fn PrintColorResults(results []telemetry.Result) {
	for _, r := range results {
		statusColor := ColorGreen
		if r.Status != telemetry.StatusUp {
			statusColor = ColorRed
		}

		fmt.Printf("[%s] Target: %s | Layer: %s | Probe: %s | Status: %s%s%s | Latency: %v\n",
			r.Timestamp.Format("15:04:05"),
			r.Target,
			r.Layer,
			r.Probe,
			statusColor, r.Status, ColorReset,
			r.Latency,
		)
	}
}
