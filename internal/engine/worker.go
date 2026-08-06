package engine

import (
	"context"

	"github.com/darnellwashington94-art/omniping/internal/layers/l3_network"
	"github.com/darnellwashington94-art/omniping/pkg/telemetry"
)

type Worker struct {
	ID int
}

func NewWorker(id int) *Worker {
	return &Worker{ID: id}
}

func (w *Worker) ExecuteProbe(ctx context.Context, target string) telemetry.Result {
	return l3_network.ProbeICMP(ctx, target)
}
