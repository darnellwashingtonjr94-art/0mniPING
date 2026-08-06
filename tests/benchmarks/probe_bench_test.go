package benchmarks

import (
	"context"
	"testing"

	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l3_network"
)

func BenchmarkProbeICMP(b *testing.B) {
	ctx := context.Background()
	target := "127.0.0.1"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l3_network.ProbeICMP(ctx, target)
	}
}
