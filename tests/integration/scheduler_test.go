package integration

import (
	"context"
	"testing"
	"time"

	"github.com/username/omniping/internal/engine"
	"github.com/username/omniping/pkg/telemetry"
)

func TestSchedulerTick(t *testing.T) {
	runner := engine.NewRunner([]string{"127.0.0.1"})
	tickCount := 0

	scheduler := engine.NewScheduler(runner, 50*time.Millisecond, func(results []telemetry.Result) {
		tickCount++
	})

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
	defer cancel()

	go scheduler.Start(ctx)

	<-ctx.Done()
	if tickCount == 0 {
		t.Fatal("expected at least one scheduler tick, got 0")
	}
}
