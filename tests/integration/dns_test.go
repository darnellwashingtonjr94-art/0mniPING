package integration

import (
	"context"
	"testing"

	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/dns"
	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func TestIntegrationDNS(t *testing.T) {
	ctx := context.Background()
	res := dns.ProbeDNSResolver(ctx, "localhost")

	if res.Status != telemetry.StatusUp {
		t.Fatalf("expected status UP for localhost DNS lookup, got DOWN: %s", res.Error)
	}
}
