package exporter

import (
	"fmt"

	"github.com/darnellwashintonjr94-art/omniping/pkg/telemetry"
)

func PrintResults(results []telemetry.Result) {
	for _, r := range results {
		fmt.Printf("[%s] Target: %-15s Layer: %s Probe: %s Status: %s Latency: %v Error: %s\n",
			r.Timestamp.Format("15:04:05"), r.Target, r.Layer, r.Probe, r.Status, r.Latency, r.Error)
	}
}
