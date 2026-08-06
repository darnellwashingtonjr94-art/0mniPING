package engine

import (
	"context"
	"sync"

	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l3_network"
	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

type WorkerPool struct {
	concurrency int
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{concurrency: concurrency}
}

func (p *WorkerPool) Run(ctx context.Context, targets []string) []telemetry.Result {
	results := make([]telemetry.Result, len(targets))
	var wg sync.WaitGroup
	sem := make(chan struct{}, p.concurrency)

	for i, t := range targets {
		wg.Add(1)
		go func(idx int, target string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			results[idx] = l3_network.ProbeICMP(ctx, target)
		}(i, t)
	}

	wg.Wait()
	return results
}
