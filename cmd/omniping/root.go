package main

import (
	"context"
	"log"

	"github.com/darnellwashingtonjr94-art/omniping/internal/engine"
	"github.com/darnellwashingtonjr94-art/omniping/internal/exporter"
)

func ExecuteRun(cli *CLIConfig) {
	ctx := context.Background()

	targets := []string{cli.Target}
	runner := engine.NewRunner(targets)

	if cli.Metrics != "" {
		promServer := exporter.NewPrometheusServer(cli.Metrics)
		go func() {
			log.Printf("Starting Prometheus metrics server on %s", cli.Metrics)
			if err := promServer.Start(); err != nil {
				log.Fatalf("Metrics server failed: %v", err)
			}
		}()
		results := runner.RunAll(ctx)
		promServer.Update(results)
	}

	results := runner.RunAll(ctx)
	_ = RenderOutput(results, cli.JSONLog)
}
