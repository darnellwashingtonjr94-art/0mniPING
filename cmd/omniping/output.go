package main

import (
	"github.com/darnellwashingtonjr94-art/omniping/internal/exporter"
	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func RenderOutput(results []telemetry.Result, jsonLogPath string) error {
	exporter.PrintResults(results)

	if jsonLogPath != "" {
		jsonExp := exporter.NewJSONExporter(jsonLogPath)
		return jsonExp.Export(results)
	}

	return nil
}
