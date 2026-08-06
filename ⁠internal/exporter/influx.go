package exporter

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

type InfluxExporter struct {
	endpoint   string
	httpClient *http.Client
}

func NewInfluxExporter(endpoint string) *InfluxExporter {
	return &InfluxExporter{
		endpoint:   endpoint,
		httpClient: &http.Client{Timeout: 3 * time.Second},
	}
}

func (i *InfluxExporter) Export(results []telemetry.Result) error {
	var sb strings.Builder
	for _, res := range results {
		upVal := 0
		if res.Status == telemetry.StatusUp {
			upVal = 1
		}
		line := fmt.Sprintf("omniping,target=%s,layer=%s,probe=%s status=%di,latency_ns=%di %d\n",
			res.Target, res.Layer, res.Probe, upVal, res.Latency.Nanoseconds(), res.Timestamp.UnixNano())
		sb.WriteString(line)
	}
	return nil
}
