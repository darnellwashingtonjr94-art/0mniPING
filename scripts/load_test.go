package main

import (
	"context"
	"fmt"
	"time"

	"github.com/darnellWashingtonjr94-art/omniping/internal/engine"
)

func RunLoadSimulation() {
	targets := []string{"127.0.0.1", "1.1.1.1", "8.8.8.8"}
	runner := engine.NewRunner(targets)
	ctx := context.Background()

	start := time.Now()
	for i := 0; i < 1000; i++ {
		_ = runner.RunAll(ctx)
	}
	duration := time.Since(start)
	fmt.Printf("Completed 1000 iteration load simulation across %d targets in %v\n", len(targets), duration)
}
