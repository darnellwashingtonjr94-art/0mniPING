package exporter

import (
	"fmt"
	"log/syslog"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

type SyslogExporter struct {
	writer *syslog.Writer
}

func NewSyslogExporter() (*SyslogExporter, error) {
	w, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "omniping")
	if err != nil {
		return nil, err
	}
	return &SyslogExporter{writer: w}, nil
}

func (s *SyslogExporter) Export(results []telemetry.Result) error {
	for _, res := range results {
		msg := fmt.Sprintf("Target: %s Layer: %s Probe: %s Status: %s Latency: %v",
			res.Target, res.Layer, res.Probe, res.Status, res.Latency)
		if res.Status == telemetry.StatusUp {
			s.writer.Info(msg)
		} else {
			s.writer.Err(msg)
		}
	}
	return nil
}
