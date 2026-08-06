package integration

import (
	"context"
	"testing"

	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l7_app"
	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
	"github.com/darnellwashingtonjr94-art/omniping/tests/mocks"
)

func TestIntegrationHTTPHealth(t *testing.T) {
	srv := mocks.NewTestHTTPServer()
	defer srv.Close()

	ctx := context.Background()
	res := l7_app.ProbeHTTPHealth(ctx, srv.URL)

	if res.Status != telemetry.StatusUp {
		t.Fatalf("expected status UP, got DOWN with error: %s", res.Error)
	}
}
