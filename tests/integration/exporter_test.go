package integration

import (
	"os"
	"testing"
	"time"

	"github.com/username/omniping/internal/exporter"
	"github.com/darnellWashingtonjr94-art/omniping/pkg/telemetry"
)

func TestJSONExporterIntegration(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "omniping-test-*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	jExp := exporter.NewJSONExporter(tmpFile.Name())
	results := []telemetry.Result{
		{
			Timestamp: time.Now(),
			Target:    "127.0.0.1",
			Layer:     "L3",
			Probe:     "icmp",
			Status:    telemetry.StatusUp,
			Latency:   time.Millisecond,
		},
	}

	if err := jExp.Export(results); err != nil {
		t.Fatalf("failed to export JSON results: %v", err)
	}
}
