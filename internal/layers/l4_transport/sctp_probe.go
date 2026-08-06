package l4_transport

import (
	"context"
	"time"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func ProbeSCTP(ctx context.Context, target string) telemetry.Result {
	start := time.Now()
	res := telemetry.Result{
		Timestamp: start,
		Target:    target,
		Layer:     "L4",
		Probe:     "sctp_probe",
	}

	// Simulation for SCTP transport layer association probing
	select {
	case <-ctx.Done():
		res.Status = telemetry.StatusDown
		res.Error = ctx.Err().Error()
		return res
	case <-time.After(20 * time.Millisecond):
		res.Status = telemetry.StatusUp
		res.Latency = time.Since(start)
	}

	return res
}
