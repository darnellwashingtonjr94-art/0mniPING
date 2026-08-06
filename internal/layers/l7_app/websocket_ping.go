package l7_app

import (
	"context"
	"time"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func ProbeWebSocket(ctx context.Context, target string) telemetry.Result {
	start := time.Now()
	res := telemetry.Result{
		Timestamp: start,
		Target:    target,
		Layer:     "L7",
		Probe:     "websocket_ping",
	}

	// Simulation for WebSocket handshake & ping frame exchange
	select {
	case <-ctx.Done():
		res.Status = telemetry.StatusDown
		res.Error = ctx.Err().Error()
		return res
	case <-time.After(25 * time.Millisecond):
		res.Status = telemetry.StatusUp
		res.Latency = time.Since(start)
	}

	return res
}
