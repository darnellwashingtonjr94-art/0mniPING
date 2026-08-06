package telemetry

import (
	"fmt"
)

func FormatSummaryLine(r Result) string {
	return fmt.Sprintf("[%s] %s | Layer: %s | Probe: %s | Status: %s | Latency: %v",
		r.Timestamp.Format("15:04:05.000"),
		r.Target,
		r.Layer,
		r.Probe,
		r.Status,
		r.Latency,
	)
}
