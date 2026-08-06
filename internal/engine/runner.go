package engine

import (
	"context"
	"sync"

	"github.com/username/omniping/internal/layers/l3_network"
	"github.com/username/omniping/pkg/telemetry"
)

type Runner struct {
	Targets []string
}

func NewRunner(targets []string) *Runner {
	return &Runner{Targets: targets}
}

func (r *Runner) RunAll(ctx context.Context) []telemetry.Result {
	var wg sync.WaitGroup
	results := make([]telemetry.Result, len(r.Targets))
	var mu sync.Mutex

	for i, t := range r.Targets {
		wg.Add(1)
		go func(idx int, target string) {
			defer wg.Done()
			res := l3_network.ProbeICMP(ctx, target)
			mu.Lock()
			results[idx] = res
			mu.Unlock()
		}(i, t)
	}
	wg.Wait()
	return results
}
